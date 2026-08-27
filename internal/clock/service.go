package clock

import (
	"sync"
	"time"
)

type Clock struct {
	mu  sync.RWMutex
	now func() time.Time
}

func New() *Clock { return &Clock{now: time.Now} }

func (c *Clock) Now() time.Time { c.mu.RLock(); f := c.now; c.mu.RUnlock(); return f().UTC() }

func (c *Clock) Set(f func() time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if f == nil {
		c.now = time.Now
	} else {
		c.now = f
	}
}

func (c *Clock) Since(t time.Time) time.Duration { return c.Now().Sub(t) }

func (c *Clock) Within(t time.Time, d time.Duration) bool { return c.Since(t) >= 0 && c.Since(t) <= d }
