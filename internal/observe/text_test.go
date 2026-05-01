package observe

import (
	"strings"
	"testing"
)

func TestRenderPrometheusIncludesCoreMetrics(t *testing.T) {
	out := RenderPrometheus(Snapshot{
		UploadRateBPS:      50_000_000,
		DownloadRateBPS:    95_000_000,
		FastLatencySeconds: 0.008,
		SlowLatencySeconds: 0.010,
		ProbeHealthy:       true,
		PriorityActive:     true,
	})

	for _, want := range []string{
		"aura_shaper_rate_bits_per_second",
		"aura_latency_fast_ewma_seconds",
		"aura_probe_health 1",
		"aura_priority_rule_active 1",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected output to contain %q:\n%s", want, out)
		}
	}
}
