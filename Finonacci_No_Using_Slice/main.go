package main

import "fmt"

//FIBONACCI: Fibonacci sequence is a sequence in which each number is the sum of the two preceding ones
//uSING SLICE AND APPEND OPERATION

//If we want the element at 6th number
func main() {

	n := 6
	fibo := []int{0, 1}

	//Lets use while loop , we use a for loop as while loop in golang
	for len(fibo) <= n {

		fibo = append(fibo, fibo[len(fibo)-1]+fibo[len(fibo)-2])
	}

	fmt.Println(fibo)
	fmt.Printf("The last element of Fibonach=ci series at postion %d is %v", n, fibo[n-1])
}
