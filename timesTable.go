package main

import "fmt"

/*
timesTable is a function that takes in two numbers as parameters,
row and col, which represent the number of rows and columns for the multiplication table respectively.
It prints a multiplication table based on the provided row and col values.
*/
func timesTable(row int, col int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Printf("%4d x %v = %-4d", i, j, i*j)
		}
		fmt.Println()
	}
}

//prints the 10-times table from 1-12
func main() {
	timesTable(12, 10)
}
