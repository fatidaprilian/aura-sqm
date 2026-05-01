package filter

import "testing"

func TestEWMAStartsWithFirstSample(t *testing.T) {
	ewma := NewEWMA(0.3)

	got := ewma.Add(0.010)
	if got != 0.010 {
		t.Fatalf("expected first sample, got %f", got)
	}
	if !ewma.Ready() {
		t.Fatal("expected EWMA to be ready after first sample")
	}
}

func TestEWMAAppliesAlpha(t *testing.T) {
	ewma := NewEWMA(0.5)
	ewma.Add(10)

	got := ewma.Add(20)
	if got != 15 {
		t.Fatalf("expected weighted value 15, got %f", got)
	}
}

func TestRejectOutlier(t *testing.T) {
	if !RejectOutlier(30, 10, 1.5) {
		t.Fatal("expected sample to be rejected")
	}
	if RejectOutlier(12, 10, 1.5) {
		t.Fatal("expected sample to be accepted")
	}
}
