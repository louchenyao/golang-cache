package goc

type Error string
type MissError struct{}

func (e MissError) Error() string {
	return "Cache miss!"
}
func (e Error) Error() string {
	return string(e)
}

type backendInterface interface {
	set(key string, val interface{})
	get(key string) (interface{}, error)
	flush(key string)
}

type Cache struct {
	countSet, countGet, countSucc, countMiss int64
	c                                        backendInterface
}

func NewCache(backend string, maxCap int) (*Cache, error) {
	c := &Cache{}

	switch backend {
	case "fake":
		c.c = newFakeCache(maxCap)
	default:
		return nil, Error("Unknow backend: " + backend)
	}

	return c, nil
}

func (c *Cache) Set(key string, val interface{}) {
	c.countSet++
	c.c.set(key, val)
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.countGet++
	val, err := c.c.get(key)

	switch err.(type) {
	case MissError:
		c.countMiss++
	case nil:
		c.countSucc++
	}

	return val, err
}
