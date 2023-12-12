package maths

import (
	"fmt"
)

func fmain() {

	var (
		names       []string
		ages        []int
		occupations []string
		states      []string
	)
	for i := 0; i < 4; i++ {
		fmt.Printf("enter your detail #%d:\n", i+1)
		var name, occupation, state string
		var age int

		fmt.Println("Enter your name")
		fmt.Scanln(&name)

		fmt.Println("Enter your age")
		fmt.Scanln(&age)

		fmt.Println("Enter your occupation")
		fmt.Scanln(&occupation)

		fmt.Println("Enter your state")
		fmt.Scanln(&states)

		names = append(names, name)
		ages = append(ages, age)
		occupations = append(occupations, occupation)
		states = append(states, state)

		fmt.Printf("list of names: %s\n", names)
		fmt.Printf("list of age: %d\n", age)
		fmt.Printf("list of occupation: %s\n", occupation)
		fmt.Printf("list of state: %s\n", states)
	}
	// fmt.Println("Collected data")

	// for i := 0; i < 4; i++ {
	// fmt.Printf("person #%d -Name: %s,age:%d, Occupation: %, states: %s\n", i+1, names[i], ages[i], occupations[i], states[i])
	// }
}
