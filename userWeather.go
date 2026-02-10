package main

import(
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
	//	"time"
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
			learnNewAreas()
			return
		default:
			fmt.Println("Invalid Number. Try again")
		}
	}
}
func findWeather(){
	fmt.Println("Finding Weather...")
	fmt.Print("Which location would you like to find the weather for?")
	reader := bufio.NewReader(os.Stdin)
	const WEATHER_API = 

	var location string
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error has occurred. Please try again")
			continue
		}
		location = strings.TrimSpace(input)
		if location == ""{
			fmt.Println("Location Can't Be Empty")
			continue
		}
		if _, err := strconv.Atoi(location); err == nil{
			fmt.Println("Please enter non-numeric location")
			continue
		}
			
		fmt.Println(location)
	//	n, err := fmt.Scan(&location)
	//	if n != 1 || err != nil {
	//		fmt.Println("Invalid Input Type")
	//		continue
	//	}
		// Pass the location into the latitude and longitude finder and then pass that information to the weather api using that information
		//TODO: We need to determine whether the location is invalid somehow..
		//switch location {
		//	case 
		//}


	}



}

func learnNewAreas(){
	fmt.Println("Learning New Areas...")

}
