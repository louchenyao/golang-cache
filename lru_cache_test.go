package goc

import (
	"strconv"
	"testing"
)

func TestBasicFunc(t *testing.T) {
	c, err := NewCache("lru", 1024)
	if err != nil {
		t.Fatal(err)
	}

	kv := map[string]interface{}{
		"1":  "Hello",
		"hi": "Hello",
		"2":  123,
	}

	for k, v := range kv {
		c.Set(k, v)
	}

	for k, v := range kv {
		ret, found := c.Get(k)
		if !found {
			t.Fatal()
		}
		if ret != v {
			t.Fatalf("Expect %s, but it's %s", v, ret)
		}
	}

	// Test not found
	v, found := c.Get("miss")
	if v != nil || found {
		t.Fatalf("Expect 'found' be false.")
	}
}

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
