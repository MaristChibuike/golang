package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var initialScore int = 100
var point int = 20
var guess string
var correctCount int
var wrongCount int
var correct string = "Wow! Correct"
var wrong string = "You guessed wronged!"
var count int
var randomNumList []int
var guessNumList []int

func main() {
	fmt.Println("Welcome to the guess game! You have", initialScore, "Ekido")
	fmt.Println("Each corrrect or wrong guess adds or removes", point, "from your Ekido respectively")
	fmt.Println("Guess a Number between 1 to 10 or 0 t0 quit.")

	for count = 0; initialScore > 0; count++ {
		var randomNumber = rand.Intn(11)

		fmt.Print("(", count+1, ") ", "Guess Number >>> ")
		fmt.Scan(&guess)
		userInput, err := strconv.Atoi(guess)

		if err != nil || userInput < 0 || userInput > 10 { // Define number range range
			fmt.Println("Enter a valid number")
			continue
		}

		if guess == "q" || guess == "Q" {
			fmt.Println("Quiting....")
			break /*"q" || guess == "Q"*/
		}
		if (userInput) == randomNumber {
			initialScore += point
			fmt.Println(correct, "Your Ekido is now:", initialScore)
			correctCount++
		} else {
			initialScore -= point
			fmt.Println(wrong, "Your Ekido is now:", initialScore)
			wrongCount++
		}
		randomNumList = append(randomNumList, randomNumber)
		guessNumList = append(guessNumList, userInput)
	}
	fmt.Printf("\nYour have guessed %d times\nCorrect guess - %d\nWrong guess - %d\nYour Final Credit is:%d\n", count, correctCount, wrongCount, initialScore)
	fmt.Println("Random Numbers  -", randomNumList)
	fmt.Println("Guessed Numbers -", guessNumList)
}
