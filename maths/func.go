package maths

import (
	"fmt"
)

// Write a function in Go that takes two integers as parameters and returns their sum.
func sum(num1 int, num2 int) int {
	return num1 + num2
}

// Create a function that accepts a variable number of strings and concatenates them into a single string.
func name(firstName string, lastName string, otherNames string) string {
	fullName := fmt.Sprintf("%v %v %v", firstName, lastName, otherNames)
	return fullName
}
func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(name("Marist", "Chibuike", "O"))
}
