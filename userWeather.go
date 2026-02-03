package main

import(
	"fmt"
)

func main(){
	var userNumChoice int
	fmt.Println("Welcome to the Weather Terminal \n What would you like to do")
	fmt.Println("1. Find the current weather for a specific area?")
	fmt.Println("2. Learn about new areas around the world to find the weather")

	for {
		fmt.Print("Input a number: ")
		n, err := fmt.Scan(&userNumChoice)

		// Printing error messages if any from scan input
		//In this case an N val of 1 means that the method had ran successfully
		//Error value of <nil> is quite self explanatory
		fmt.Println("N val: ", n, "Err Val:", err)

		switch userNumChoice {
		case 1:
			findWeather()
			return
		case 2:
			learnNewAreas()
			return
		case 3:
			fmt.Print("Invalid Input. Try again")
		}

	}




}


func findWeather(){
	fmt.Println("Finding Weather...")


}

func learnNewAreas(){
	fmt.Println("Learning New Areas...")

}
