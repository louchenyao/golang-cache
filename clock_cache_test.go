package goc

import (
	"strconv"
	"testing"
)

func TestClockReplacement(t *testing.T) {
	c := newClockCache(10)
	for i := 0; i < 10; i++ {
		c.set(strconv.Itoa(i), i)
	}

	// touch even
	for _j := 0; _j <= 100; _j++ {
		for i := 0; i < 10; i += 2 {
			_, _ = c.get(strconv.Itoa(i))
		}
	}

	// touch more except 1
	for i := 0; i <= 100000; i++ {
		if i%10 == 1 {
			continue
		}
		_, _ = c.get(strconv.Itoa(i % 10))
	}

	for i := 0; i < 10; i++ {
		t.Logf("c.items[%d].count = %d", i, c.items[i].count)
	}

	// refill someting
	for i := 20; i < 25; i++ {
		c.set(strconv.Itoa(i), i)
	}

	// even sould still in
	for i := 0; i < 10; i += 2 {
		val, ok := c.get(strconv.Itoa(i))
		if !ok {
			t.Fatalf("%d not in cache", i)
		}
		if val != i {
			t.Fatalf("Excpect %d", i)
		}
	}

}
