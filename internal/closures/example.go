package closures

import "fmt"

// closures

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func callingAClosure() {
	f, g := fib(), fib()
	fmt.Println(f(), g())
}
