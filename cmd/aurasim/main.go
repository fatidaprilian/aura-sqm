package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fatidaprilian/aura-sqm/internal/config"
	auraruntime "github.com/fatidaprilian/aura-sqm/internal/runtime"
	"github.com/fatidaprilian/aura-sqm/internal/observe"
	"github.com/fatidaprilian/aura-sqm/internal/probe"
	"github.com/fatidaprilian/aura-sqm/internal/shaper"
)

func main() {
	configPath := flag.String("config", "config/example.json", "path to Aura-SQM JSON configuration")
	ticks := flag.Int("ticks", 120, "number of simulation ticks to run")
	serveMetrics := flag.Bool("serve-metrics", false, "serve /metrics while simulation runs")
	flag.Parse()

	cfg, err := config.LoadFile(*configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	controller := shaper.NewMemoryController(shaper.Rates{
		UploadBPS:   cfg.Shaper.UploadCeilingMbps * 1_000_000,
		DownloadBPS: cfg.Shaper.DownloadCeilingMbps * 1_000_000,
	})
	source := &probe.ScriptedSource{
		ReflectorID:   "sim-cloudflare",
		Protocol:      "icmp",
		BaseLatency:   cfg.Control.TargetLatencyMS / 1000,
		BufferLatency: 0.025,
		SpikeEvery:    37,
		SpikeLatency:  0.080,
	}
	engine := auraruntime.NewEngine(cfg, source, controller)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if *serveMetrics {
		server := observe.NewServer(cfg.Observability.MetricsListen, engine)
		go func() {
			if err := server.ListenAndServe(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				cancel()
			}
		}()
		defer server.Shutdown(context.Background())
	}

	for i := 0; i < *ticks; i++ {
		if err := engine.Tick(ctx); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if i%10 == 0 || i == *ticks-1 {
			fmt.Print(observe.RenderText(engine.Snapshot()))
		}
		time.Sleep(time.Duration(cfg.Control.LoopIntervalMS) * time.Millisecond)
	}
}
