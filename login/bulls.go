package login

import "fmt"

func Bulls(bulidings string) {
	switch bulidings {
	case "flateBuilding":
		fmt.Println("The are all occupied")
	case "duplexBuilding":
		fmt.Println("Can you pay the amount required?")
	case "semiduplexBuilding":
		fmt.Println("The offical is 250k?")
	case "bungaloBuilding":
		fmt.Println("The house is for the rich?")
	default:
		fmt.Println("You are in a wrong app?")
	}
}
