package main

import "fmt"

// print numbers 1 to n backward
func print(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Println(n - 1)
	return print(n - 1)
}

func main() {
	fmt.Println(print(5))
}
