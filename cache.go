package goc

import "sync"

type Error string
type MissError struct{}

func (e Error) Error() string {
	return string(e)
}

type backendInterface interface {
	set(key string, val interface{})
	get(key string) (interface{}, bool)
	flush(key string)
}

type Cache struct {
	countSet, countGet, countMiss int64
	c                             backendInterface
	lock                          sync.RWMutex
}

func NewCache(backend string, maxCap int) (*Cache, error) {
	c := &Cache{}

	switch backend {
	case "fake":
		c.c = newFakeCache(maxCap)
	case "lru":
		c.c = newLruCache(maxCap)
	case "clock":
		c.c = newClockCache(maxCap)
	default:
		return nil, Error("Unknow backend: " + backend)
	}

	return c, nil
}

func (c *Cache) Set(key string, val interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.countSet++
	c.c.set(key, val)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.countGet++
	val, found := c.c.get(key)

	if found {
		c.countMiss++
	}

	return val, found
}
