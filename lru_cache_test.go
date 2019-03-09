package goc

import (
	"strconv"
	"testing"
)

func TestGC(t *testing.T) {
	c := newLruCache(1024)

	for i := 0; i < 1024*3+5; i++ {
		c.set(strconv.Itoa(i), i+1)
	}

	v, ok := c.get("3")
	if ok {
		t.Fatal("key 3 should already be deleted")
	}
	v, ok = c.get(strconv.Itoa(1024*3 - 2))
	if !ok || v.(int) != 1024*3-1 {
		t.Fatal("key 1024 * 3 - 2 miss")
	}

	cnt := 0
	for i := 0; i < 1024*3+5; i++ {
		v, ok = c.get(strconv.Itoa(i))
		//fmt.Println(v, ok)
		if ok {
			cnt += 1
		}
	}
	if cnt > 2*1024+1024 || c.cap != cnt {
		t.Fatal("gc not worked")
	}
}
