<div style="display: flex; justify-content: center;">
  <img src="./img/banner.png"/>
</div>

# Chinh phuc Tour of Go

## Packages, variables, and functions

## Flow control statements: for, if, else, switch and defer
Learn how to control the flow of your code with conditionals, loops, switches and defers.

### Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number `x`, we want to find the number `z` for which `z²` is most nearly `x`.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

```z -= (z*z - x) / (2*z)```

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

> Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:
```
z := 1.0
z := float64(1)
```
Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?

> (Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)

```Go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:= 1.0
	for i:=0;i<10;i++ {
		z-=(z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
```

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
## Methods and interfaces
Learn how to define methods on types, how to declare interfaces, and how to put everything together.

### Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

```Go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

### Exercise: Errors
Copy your `Sqrt` function from the earlier exercise and modify it to return an `error` value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

```type ErrNegativeSqrt float64```
and make it an error by giving it a

```func (e ErrNegativeSqrt) Error() string```
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

> *Note*: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an `ErrNegativeSqrt` value when given a negative number.

```Go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z:=1.0
		for i:=0;i<10;i++ {
			z-=(z*z - x) / (2*z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(-4))
}
```

# References
[Go doc](https://golang.org/doc/)

[A tour of Go](https://tour.golang.org/list)
