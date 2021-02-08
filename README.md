<div style="display: flex; justify-content: center;">
  <img src="./img/banner.png"/>
</div>

# Chinh phuc Tour of Go

## Packages, variables, and functions

## Flow control statements: for, if, else, switch and defer

## More types: structs, slices, and maps.

### Exercise: Slices

Implement **Pic**. It should *return a slice of length dy*, each element of which is *a slice of dx 8-bit unsigned integers*. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

```Go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rs:= make([][]uint8, dy)
	for y:= range rs {
		rs[y]=make([]uint8, dx)
		for x:= range rs[y]{
			rs[y][x] = uint8(x^y)
		}
	}
	return rs
}

func main() {
	pic.Show(Pic)
}
```

### Exercise: Maps

Implement **WordCount**. It should return *a map of the counts of each “word” in the string s*. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find [strings.Fields](https://golang.org/pkg/strings/#Fields) helpful.

> first version
```Go
package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	var arr = strings.Fields(s)
	m := make(map[string]int)

	for v := range arr {
		elem, ok := m[arr[v]]
		if ok {
			m[arr[v]] = elem + 1
		} else {
			m[arr[v]] = 1
		}
	}
	return m
}

func main() {
	for i, v := range strings.Fields("Yen lun") {
		fmt.Println(i,v)
	}
}
```

> latest version
```Go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m:=make(map[string]int)
	for _,v:= range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

### Exercise: Fibonacci closure
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

```Go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	v0, v1:= 0, 1
	return func() int {
		v:= v0
		v0, v1 = v1, v1+v
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

# References
[Go doc](https://golang.org/doc/)

[A tour of Go](https://tour.golang.org/list)
