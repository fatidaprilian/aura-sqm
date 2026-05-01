package runtime

import (
	"context"
	"sync"
	"time"

	"github.com/fatidaprilian/aura-sqm/internal/config"
	"github.com/fatidaprilian/aura-sqm/internal/control"
	"github.com/fatidaprilian/aura-sqm/internal/filter"
	"github.com/fatidaprilian/aura-sqm/internal/observe"
	"github.com/fatidaprilian/aura-sqm/internal/probe"
	"github.com/fatidaprilian/aura-sqm/internal/shaper"
)

type Engine struct {
	cfg       config.Config
	source    probe.Source
	shaper    shaper.Controller
	governor  *control.PID
	fastEWMA  filter.EWMA
	slowEWMA  filter.EWMA
	mu        sync.RWMutex
	snapshot  observe.Snapshot
	tickTotal uint64
}

func NewEngine(cfg config.Config, source probe.Source, controller shaper.Controller) *Engine {
	initial := controller.Current()
	return &Engine{
		cfg:    cfg,
		source: source,
		shaper: controller,
		governor: control.NewPID(control.PIDConfig{
			KP:              cfg.Control.KP,
			KI:              cfg.Control.KI,
			KD:              cfg.Control.KD,
			IntegralMin:     cfg.Control.IntegralMin,
			IntegralMax:     cfg.Control.IntegralMax,
			MaxRateDeltaBPS: cfg.Control.MaxRateDeltaMbps * 1_000_000,
		}),
		fastEWMA: filter.NewEWMA(cfg.Probe.FastEWMAAlpha),
		slowEWMA: filter.NewEWMA(cfg.Probe.SlowEWMAAlpha),
		snapshot: observe.Snapshot{
			UploadRateBPS:   initial.UploadBPS,
			DownloadRateBPS: initial.DownloadBPS,
			ProbeHealthy:    true,
			PriorityActive:  cfg.Priority.Enabled,
		},
	}
}

func (e *Engine) Run(ctx context.Context) error {
	interval := time.Duration(e.cfg.Control.LoopIntervalMS) * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := e.Tick(ctx); err != nil {
				return err
			}
		}
	}
}

func (e *Engine) Tick(ctx context.Context) error {
	sample, err := e.source.Next(ctx)
	if err != nil {
		return err
	}

	current := e.shaper.Current()
	fastLatency := e.fastEWMA.Value()
	slowLatency := e.slowEWMA.Value()

	if sample.Healthy {
		if !e.slowEWMA.Ready() || !filter.RejectOutlier(sample.LatencySeconds, e.slowEWMA.Value(), e.cfg.Probe.OutlierThreshold) {
			fastLatency = e.fastEWMA.Add(sample.LatencySeconds)
			slowLatency = e.slowEWMA.Add(sample.LatencySeconds)
		}
	}

	decision := e.governor.Step(control.Input{
		TargetLatencySeconds:  e.cfg.Control.TargetLatencyMS / 1000,
		CurrentLatencySeconds: fastLatency,
		CurrentRateBPS:        current.UploadBPS,
		FloorBPS:              e.cfg.Shaper.UploadFloorMbps * 1_000_000,
		CeilingBPS:            e.cfg.Shaper.UploadCeilingMbps * 1_000_000,
		DeltaSeconds:          float64(e.cfg.Control.LoopIntervalMS) / 1000,
		ProbeHealthy:          sample.Healthy,
	})

	next := shaper.Rates{
		UploadBPS:   decision.NextRateBPS,
		DownloadBPS: current.DownloadBPS,
	}
	if err := e.shaper.Apply(ctx, next); err != nil {
		return err
	}

	e.mu.Lock()
	e.tickTotal++
	e.snapshot = observe.Snapshot{
		UploadRateBPS:      next.UploadBPS,
		DownloadRateBPS:    next.DownloadBPS,
		FastLatencySeconds: fastLatency,
		SlowLatencySeconds: slowLatency,
		ProbeHealthy:       sample.Healthy,
		FallbackActive:     decision.FallbackActive,
		PriorityActive:     e.cfg.Priority.Enabled,
		ControlError:       decision.Error,
		ControlIntegral:    decision.Integral,
		ControlDerivative:  decision.Derivative,
		TickTotal:          e.tickTotal,
	}
	e.mu.Unlock()

	return nil
}

func (e *Engine) Snapshot() observe.Snapshot {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.snapshot
}
