package goc

import "sort"

type item struct {
	key     string
	val     interface{}
	lastVis int64
}

type lruCache struct {
	cap, maxCap int
	logicClock  int64
	m           map[string]item
}

func (c *lruCache) set(key string, val interface{}) {
	c.cap++
	c.logicClock++
	c.m[key] = item{key, val, c.logicClock}

	// Discards the least recently used item in lazy method.
	// The amortized complexity for each set opration is O(logn). n is maxCap.
	// Because we do gc after O(n) operations, and each gc will cost O(n*logn).
	if c.cap > 2*c.maxCap+1024 {
		c.gc()
	}
}

func (c *lruCache) get(key string) (interface{}, bool) {
	if itemVal, ok := c.m[key]; ok {
		// Updating the last visting time
		c.logicClock++
		itemVal.lastVis = c.logicClock
		c.m[key] = itemVal

		return itemVal.val, true
	}
	return nil, false
}

func (c *lruCache) flush(key string) {
	c.cap--
	delete(c.m, key)
}

type byLastVisDesc []item

func (a byLastVisDesc) Len() int           { return len(a) }
func (a byLastVisDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLastVisDesc) Less(i, j int) bool { return a[i].lastVis > a[j].lastVis }

func (c *lruCache) gc() {
	var items []item
	for _, v := range c.m {
		items = append(items, v)
	}

	sort.Sort(byLastVisDesc(items))

	c.m = make(map[string]item)
	for i := 0; i < len(items) && i < c.maxCap; i++ {
		c.m[items[i].key] = items[i]
	}
	c.cap = len(c.m)
}

func newLruCache(maxCap int) *lruCache {
	c := &lruCache{}
	c.maxCap = maxCap
	c.m = make(map[string]item)
	return c
}
