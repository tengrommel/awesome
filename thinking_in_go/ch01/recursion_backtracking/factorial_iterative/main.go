package main

import "fmt"

// calculates factorial of a positive integer
func factorial(n int) int {
	if n == 0 { // base cases: fact of 0
		return 1
	}
	return n * factorial(n-1)
}

var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers

// factorial iterative algorithm
func factorialIterative(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // mismatched types int64 and int
		}
	}
	return factVal /* return from function */
}

func main() {
	fmt.Println(factorialIterative(5))
}
