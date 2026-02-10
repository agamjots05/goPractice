package main

import(
	"fmt"
//	"bufio"
	
)


func main(){
	var userNumChoice int
	fmt.Println("Welcome to the Weather Terminal \n What would you like to do")
	fmt.Println("1. Find the current weather for a specific area?")
	fmt.Println("2. Learn about new areas around the world to find the weather")

	for {
		fmt.Print("Input a number: ")
		n, err := fmt.Scan(&userNumChoice)
		if err != nil || n != 1{
			fmt.Println("Invalid Input Type")
			continue
		}
		switch userNumChoice {
		case 1:
			findWeather()
			return
		case 2:
//			learnNewAreas()
			return
		default:
			fmt.Println("Invalid Number. Try again")
		}
	}
}

