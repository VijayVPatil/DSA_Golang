package main

import "fmt"

func main() {

	n := 6
	for i := 0; i < n; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {

	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
