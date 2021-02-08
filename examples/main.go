package main

import (
	"fmt"
	"time"
)

// fibonacci with loop
func fib(n uint) uint {
	a := uint(0)
	b := uint(1)

	for i := uint(0); i < n; i++ {
		a = a + b
		// fmt.Printf("%d : %d \n", i, a)
		a, b = b, a
	}
	return a
}

// fibnacci with dynamic
func fib2(n uint) uint {
	fn := make([]uint, n+1)
	for i := uint(0); i < n; i++ {
		var f = uint(0)
		if i <= 2 {
			f = uint(1)
		} else {
			f = fn[i-1] + fn[i-2]
		}
		// fmt.Printf("%d : %d \n", i, f)
		fn[i] = f
	}
	return fn[n]
}

// fibonacci with recursion
func fib3(n uint) uint {
	if n <= 1 {
		return uint(1)
	}
	return fib3(n-1) + fib3(n-2)
}

func main() {
	n := uint(13)
	start := time.Now()
	fib(n)
	elapsed := time.Since(start)

	start1 := time.Now()
	fib2(n)
	elapsed1 := time.Since(start1)

	start3 := time.Now()
	fib3(n)
	elapsed3 := time.Since(start3)

	fmt.Printf("time elapsed is %s \n", elapsed)
	fmt.Printf("time elapsed1 is %s \n", elapsed1)
	fmt.Printf("time elapsed2 is %s \n", elapsed3)

}
