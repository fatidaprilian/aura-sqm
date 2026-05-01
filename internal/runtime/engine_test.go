package runtime

import (
	"context"
	"testing"

	"github.com/fatidaprilian/aura-sqm/internal/config"
	"github.com/fatidaprilian/aura-sqm/internal/probe"
	"github.com/fatidaprilian/aura-sqm/internal/shaper"
)

func TestEngineTickUpdatesSnapshot(t *testing.T) {
	cfg := testConfig()
	controller := shaper.NewMemoryController(shaper.Rates{
		UploadBPS:   cfg.Shaper.UploadCeilingMbps * 1_000_000,
		DownloadBPS: cfg.Shaper.DownloadCeilingMbps * 1_000_000,
	})
	source := &probe.ScriptedSource{
		ReflectorID:   "test",
		Protocol:      "icmp",
		BaseLatency:   0.020,
		BufferLatency: 0,
	}
	engine := NewEngine(cfg, source, controller)

	if err := engine.Tick(context.Background()); err != nil {
		t.Fatalf("tick failed: %v", err)
	}

	snapshot := engine.Snapshot()
	if snapshot.TickTotal != 1 {
		t.Fatalf("expected one tick, got %d", snapshot.TickTotal)
	}
	if snapshot.FastLatencySeconds == 0 {
		t.Fatal("expected fast latency to be populated")
	}
}

func testConfig() config.Config {
	return config.Config{
		Shaper: config.ShaperSettings{
			AutoDiscoverWAN:     true,
			UploadFloorMbps:     50,
			UploadCeilingMbps:   100,
			DownloadFloorMbps:   50,
			DownloadCeilingMbps: 100,
			CAKEDiffserv:        "diffserv4",
			CAKEIsolation:       "triple-isolate",
		},
		Control: config.ControlSettings{
			LoopIntervalMS:  50,
			TargetLatencyMS: 8,
			KP:              0.2,
			KI:              0.02,
			KD:              0.01,
			IntegralMin:     -1,
			IntegralMax:     1,
			MaxRateDeltaMbps: 5,
		},
		Probe: config.ProbeSettings{
			Reflectors: []config.Reflector{
				{ID: "cloudflare", Host: "1.1.1.1", Kind: "public", Enabled: true},
			},
			Protocols:        []string{"icmp"},
			TimeoutMS:        500,
			FastEWMAAlpha:    0.3,
			SlowEWMAAlpha:    0.1,
			OutlierThreshold: 2,
		},
		Priority: config.PrioritySettings{
			Enabled:   true,
			DeviceIP:  "192.168.10.50",
			TargetTin: "voice",
		},
		Observability: config.ObservabilitySettings{
			MetricsListen: "127.0.0.1:9108",
			LogLevel:      "info",
			TUIEnabled:    true,
		},
	}
}
