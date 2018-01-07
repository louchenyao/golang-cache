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
	c, _ := goc.NewCache("fake", 233)
	c.Set(233, "Hello goc!")
	fmt.Println(c.Get(233))
	fmt.Println(c.Get(234))
}
~~~

Output:

~~~plain
Hello goc! <nil>
<nil> Cache miss!
~~~