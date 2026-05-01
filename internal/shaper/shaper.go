package shaper

import (
	"context"
	"sync"
)

type Rates struct {
	UploadBPS   float64
	DownloadBPS float64
}

type Controller interface {
	Apply(ctx context.Context, rates Rates) error
	Current() Rates
}

type MemoryController struct {
	mu    sync.RWMutex
	rates Rates
}

func NewMemoryController(initial Rates) *MemoryController {
	return &MemoryController{rates: initial}
}

func (c *MemoryController) Apply(ctx context.Context, rates Rates) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.rates = rates
	return nil
}

func (c *MemoryController) Current() Rates {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.rates
}
