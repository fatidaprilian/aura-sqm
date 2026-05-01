package control

import "testing"

func TestPIDClampsRateDelta(t *testing.T) {
	pid := NewPID(PIDConfig{
		KP:              1,
		KI:              0,
		KD:              0,
		IntegralMin:     -1,
		IntegralMax:     1,
		MaxRateDeltaBPS: 1_000_000,
	})

	decision := pid.Step(Input{
		TargetLatencySeconds:  0.020,
		CurrentLatencySeconds: 0.001,
		CurrentRateBPS:        50_000_000,
		FloorBPS:              10_000_000,
		CeilingBPS:            100_000_000,
		DeltaSeconds:          0.05,
		ProbeHealthy:          true,
	})

	if decision.NextRateBPS != 51_000_000 {
		t.Fatalf("expected clamped rate increase, got %.0f", decision.NextRateBPS)
	}
}

func TestPIDDoesNotTightenWhenProbeUnhealthy(t *testing.T) {
	pid := NewPID(PIDConfig{MaxRateDeltaBPS: 1_000_000})

	decision := pid.Step(Input{
		CurrentRateBPS: 40_000_000,
		FloorBPS:       10_000_000,
		CeilingBPS:     100_000_000,
		ProbeHealthy:   false,
	})

	if !decision.FallbackActive {
		t.Fatal("expected fallback")
	}
	if decision.NextRateBPS != 40_000_000 {
		t.Fatalf("expected current rate to be preserved, got %.0f", decision.NextRateBPS)
	}
}
