package counter

import "sync"

type Counters struct {
	mx sync.Mutex
	m  map[string]int
}

func New() *Counters {
	return &Counters{
		m: make(map[string]int),
	}
}

func (c *Counters) Load(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Counters) LoadAll() map[string]int {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.m
}

func (c *Counters) Store(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}
