# GOC (goalng-cache)

**It's in developing, the inferface is not stable now!**

The golang cache library with simple interface and ease for using!

~~~go
package main

import (
	"fmt"

	"github.com/Chenyao2333/golang-cache"
)

func main() {
	c, _ := goc.NewCache("fake", 1000)

	c.Set("hi", "Hello goc!")
	fmt.Println(c.Get("hi"))
	fmt.Println(c.Get("hello"))
}

~~~

Output:

~~~plain
Hello goc! <nil>
<nil> Cache miss!
~~~