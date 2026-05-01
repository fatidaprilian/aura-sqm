package filter

type EWMA struct {
	alpha       float64
	value       float64
	initialized bool
}

func NewEWMA(alpha float64) EWMA {
	return EWMA{alpha: alpha}
}

func (e *EWMA) Add(sample float64) float64 {
	if !e.initialized {
		e.value = sample
		e.initialized = true
		return e.value
	}

	e.value = e.alpha*sample + (1-e.alpha)*e.value
	return e.value
}

func (e EWMA) Value() float64 {
	return e.value
}

func (e EWMA) Ready() bool {
	return e.initialized
}
