package main

import "fmt"

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		// Function to print the output
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

// Function to generate all k-ary strings
func generateKAryStrings(n int, A []int, i int, k int) {
	if i == n {
		printResult(A, n)
		return
	}
	for j := 0; j < k; j++ {
		A[i] = j
		generateKAryStrings(n, A, i+1, k)
	}
}

func main() {
	var n = 4
	A := make([]int, n)
	// Print all binary strings
	generateKAryStrings(n, A, 0, 4)
}
