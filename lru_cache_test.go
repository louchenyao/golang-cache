package goc

import "testing"

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
