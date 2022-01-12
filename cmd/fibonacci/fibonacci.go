package main

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciIter(n int) int {
	if n < 2 {
		return n
	}

	a := 1
	b := 2

	for i := 3; i < n; i++ {
		a, b = b, b+a
	}

	return b
}

func main() {}
