package goc

// fakeCache only supports one key
type fakeCache struct {
	key string
	val interface{}
}

func newFakeCache(maxCap int) *fakeCache {
	return &fakeCache{}
}

func (c *fakeCache) set(key string, val interface{}) {
	c.key = key
	c.val = val
}

func (c *fakeCache) get(key string) (interface{}, bool) {
	if c.key == key {
		return c.val, true
	}
	return nil, false
}

func (c *fakeCache) flush(key string) {
	c.key = ""
	c.val = nil
}
