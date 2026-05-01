package observe

import "fmt"

type Snapshot struct {
	UploadRateBPS      float64
	DownloadRateBPS    float64
	FastLatencySeconds float64
	SlowLatencySeconds float64
	ProbeHealthy       bool
	FallbackActive     bool
	PriorityActive     bool
	ControlError       float64
	ControlIntegral    float64
	ControlDerivative  float64
	TickTotal          uint64
}

func RenderText(s Snapshot) string {
	return fmt.Sprintf(
		"upload_rate_bps=%.0f\n"+
			"download_rate_bps=%.0f\n"+
			"fast_latency_seconds=%.6f\n"+
			"slow_latency_seconds=%.6f\n"+
			"probe_healthy=%t\n"+
			"fallback_active=%t\n"+
			"priority_active=%t\n"+
			"control_error=%.6f\n"+
			"control_integral=%.6f\n"+
			"control_derivative=%.6f\n"+
			"tick_total=%d\n",
		s.UploadRateBPS,
		s.DownloadRateBPS,
		s.FastLatencySeconds,
		s.SlowLatencySeconds,
		s.ProbeHealthy,
		s.FallbackActive,
		s.PriorityActive,
		s.ControlError,
		s.ControlIntegral,
		s.ControlDerivative,
		s.TickTotal,
	)
}

func RenderPrometheus(s Snapshot) string {
	probeHealth := 0
	if s.ProbeHealthy {
		probeHealth = 1
	}
	fallback := 0
	if s.FallbackActive {
		fallback = 1
	}
	priority := 0
	if s.PriorityActive {
		priority = 1
	}

	return fmt.Sprintf(
		"# TYPE aura_shaper_rate_bits_per_second gauge\n"+
			"aura_shaper_rate_bits_per_second{direction=\"upload\"} %.0f\n"+
			"aura_shaper_rate_bits_per_second{direction=\"download\"} %.0f\n"+
			"# TYPE aura_latency_fast_ewma_seconds gauge\n"+
			"aura_latency_fast_ewma_seconds{direction=\"wan\"} %.6f\n"+
			"# TYPE aura_latency_slow_ewma_seconds gauge\n"+
			"aura_latency_slow_ewma_seconds{direction=\"wan\"} %.6f\n"+
			"# TYPE aura_probe_health gauge\n"+
			"aura_probe_health %d\n"+
			"# TYPE aura_panic_fallback_active gauge\n"+
			"aura_panic_fallback_active %d\n"+
			"# TYPE aura_priority_rule_active gauge\n"+
			"aura_priority_rule_active %d\n"+
			"# TYPE aura_control_loop_tick_total counter\n"+
			"aura_control_loop_tick_total %d\n"+
			"# TYPE aura_control_error_seconds gauge\n"+
			"aura_control_error_seconds %.6f\n",
		s.UploadRateBPS,
		s.DownloadRateBPS,
		s.FastLatencySeconds,
		s.SlowLatencySeconds,
		probeHealth,
		fallback,
		priority,
		s.TickTotal,
		s.ControlError,
	)
}
