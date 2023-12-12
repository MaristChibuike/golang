package main

import "fmt"

func factorial(n int) int {
	f := n // Holds the first value of n
	for i := 1; i < n; i++ {
		f *= (n - i)
	}
	return f
}
func main() {
	fmt.Print(factorial(5))
}
