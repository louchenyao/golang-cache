package goc

// fakeCache only supports one key
type fakeCache struct {
	key, val interface{}
}

func newFakeCache(maxCap int) *fakeCache {
	return &fakeCache{}
}

func (c *fakeCache) set(key interface{}, val interface{}) {
	c.key = key
	c.val = val
}

func (c *fakeCache) get(key interface{}) (interface{}, error) {
	if c.key == key {
		return c.val, nil
	}
	return nil, MissError{}
}

func (c *fakeCache) flush(key interface{}) {
	c.key = nil
	c.val = nil
}
