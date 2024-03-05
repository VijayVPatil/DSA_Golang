package main

import "fmt"

//FIBONACCI Series

func main() {

	fibonacci(6)

}

func fibonacci(n int) {
	previous := 0
	current := 1
	var new int
	for i := 0; i < n; i++ {

		fmt.Printf("Fibonacci of (%d) is %d\n", 0, previous)

		new = previous + current
		previous = current
		current = new

	}

}
