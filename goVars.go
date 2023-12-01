package main

import "fmt"

func main() {
	var stationName string
	var nextTrainTime int
	var isUptownTrain bool

	stationName = "Obiagu"
	nextTrainTime = 12
	isUptownTrain = false

	fmt.Println("Current Station:", stationName)
	fmt.Println("Next train:", nextTrainTime, "minutes")
	fmt.Println("Is uptown:", isUptownTrain)

	stationName = "Iwollo Central"
	nextTrainTime = 3
	isUptownTrain = true

	fmt.Println("")
	fmt.Println("Current station:", stationName)
	fmt.Println("Next Train:", nextTrainTime, "minutes")
	fmt.Println("Is Uptown:", isUptownTrain)
}
