package goc

import (
	"math"
)

type ccitem struct {
	key   string
	val   interface{}
	valid bool
	count int64
}

type clockCache struct {
	maxCap   int
	items    []ccitem
	freelist []int
	m        map[string]int
	clockP   int
}

func newClockCache(maxCap int) *clockCache {
	c := &clockCache{maxCap: maxCap, items: make([]ccitem, maxCap), freelist: make([]int, 0), m: make(map[string]int)}
	for i := 0; i < maxCap; i++ {
		c.freelist = append(c.freelist, i)
	}
	return c
}

func (c *clockCache) set(key string, val interface{}) {
	if i, ok := c.m[key]; ok {
		c.items[i].val = val
		return
	}
	if len(c.freelist) == 0 {
		c.evict()
	}
	var i int
	i, c.freelist = c.freelist[0], c.freelist[1:]
	//fmt.Printf("c.freelist len = %d", len((c.freelist)))
	c.items[i].key = key
	c.items[i].val = val
	c.items[i].count = 0
	c.items[i].valid = true
	c.m[key] = i
	//fmt.Printf("c.itmes[%d].val = %s\n", i, val)
}

func (c *clockCache) get(key string) (interface{}, bool) {
	if i, ok := c.m[key]; ok {
		//fmt.Printf("key = %s, i = %d, c.items[key].val = %s \n", key, i, c.items[i].val)
		c.items[i].count++
		return c.items[i].val, true
	}
	return nil, false
}

func (c *clockCache) flush(key string) {
	if i, ok := c.m[key]; ok {
		c.deleteItem(i)
	}
}

func (c *clockCache) wrapAdd1(a int) int {
	return (a + 1) % c.maxCap
}

func (c *clockCache) deleteItem(i int) {
	delete(c.m, c.items[i].key)
	c.items[i].valid = false
	c.items[i].count = 0
	c.freelist = append(c.freelist, i)
}

func (c *clockCache) evict() {
	// finding min
	var min int64 = math.MaxInt64
	i := c.clockP
	for _j := 0; _j < c.maxCap; _j++ {
		if c.items[i].count < min {
			min = c.items[i].count
		}
		// found count = 0 in finding
		if c.items[i].count == 0 {
			for c.clockP != i {
				c.items[c.clockP].count--
				c.clockP = c.wrapAdd1(c.clockP)
			}
			c.deleteItem(i)
			c.clockP = c.wrapAdd1(c.clockP)
			return
		}
		i = c.wrapAdd1(i)
	}

	// using min to find cache should be evicted
	for true {
		if c.items[c.clockP].count <= min {
			c.deleteItem(c.clockP)
			c.clockP = c.wrapAdd1(c.clockP)
			return
		}
		c.items[c.clockP].count -= min + 1
		c.clockP = c.wrapAdd1(c.clockP)
	}
}
