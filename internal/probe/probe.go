package probe

import (
	"context"
	"math"
)

type Sample struct {
	LatencySeconds float64
	Healthy        bool
	ReflectorID    string
	Protocol       string
}

type Source interface {
	Next(ctx context.Context) (Sample, error)
}

type ScriptedSource struct {
	ReflectorID     string
	Protocol        string
	BaseLatency     float64
	BufferLatency   float64
	SpikeEvery      int
	SpikeLatency    float64
	FailureStart    int
	FailureDuration int
	tick            int
}

func (s *ScriptedSource) Next(ctx context.Context) (Sample, error) {
	select {
	case <-ctx.Done():
		return Sample{}, ctx.Err()
	default:
	}

	s.tick++
	if s.FailureStart > 0 && s.tick >= s.FailureStart && s.tick < s.FailureStart+s.FailureDuration {
		return Sample{
			Healthy:     false,
			ReflectorID: s.ReflectorID,
			Protocol:    s.Protocol,
		}, nil
	}

	latency := s.BaseLatency + loadWave(s.tick)*s.BufferLatency
	if s.SpikeEvery > 0 && s.tick%s.SpikeEvery == 0 {
		latency += s.SpikeLatency
	}

	return Sample{
		LatencySeconds: latency,
		Healthy:        true,
		ReflectorID:    s.ReflectorID,
		Protocol:       s.Protocol,
	}, nil
}

func loadWave(tick int) float64 {
	phase := float64(tick%120) / 120
	return 0.5 + 0.5*math.Sin(phase*2*math.Pi)
}
