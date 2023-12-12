package main

import "fmt"

func main() {

	i := 5

	for i = 5; i < 0; i-- {
		i *= (i - 1)
		fmt.Print(i)
	}
	fmt.Print(i)
}
