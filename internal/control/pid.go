package control

type PIDConfig struct {
	KP              float64
	KI              float64
	KD              float64
	IntegralMin     float64
	IntegralMax     float64
	MaxRateDeltaBPS float64
}

type PID struct {
	cfg         PIDConfig
	integral    float64
	previousErr float64
	hasPrevious bool
}

type Input struct {
	TargetLatencySeconds  float64
	CurrentLatencySeconds float64
	CurrentRateBPS        float64
	FloorBPS              float64
	CeilingBPS            float64
	DeltaSeconds          float64
	ProbeHealthy          bool
}

type Decision struct {
	NextRateBPS    float64
	Error          float64
	Integral       float64
	Derivative     float64
	FallbackActive bool
}

func NewPID(cfg PIDConfig) *PID {
	return &PID{cfg: cfg}
}

func (p *PID) Step(in Input) Decision {
	if !in.ProbeHealthy {
		return Decision{
			NextRateBPS:    clamp(in.CurrentRateBPS, in.FloorBPS, in.CeilingBPS),
			FallbackActive: true,
		}
	}

	err := in.TargetLatencySeconds - in.CurrentLatencySeconds
	dt := in.DeltaSeconds
	if dt <= 0 {
		dt = 1
	}

	p.integral = clamp(p.integral+err*dt, p.cfg.IntegralMin, p.cfg.IntegralMax)

	derivative := 0.0
	if p.hasPrevious {
		derivative = (err - p.previousErr) / dt
	}
	p.previousErr = err
	p.hasPrevious = true

	controlSignal := p.cfg.KP*err + p.cfg.KI*p.integral + p.cfg.KD*derivative
	rateDelta := clamp(controlSignal*in.CeilingBPS, -p.cfg.MaxRateDeltaBPS, p.cfg.MaxRateDeltaBPS)
	nextRate := clamp(in.CurrentRateBPS+rateDelta, in.FloorBPS, in.CeilingBPS)

	return Decision{
		NextRateBPS: nextRate,
		Error:       err,
		Integral:    p.integral,
		Derivative:  derivative,
	}
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
