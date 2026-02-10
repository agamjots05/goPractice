package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func findWeather(){
	fmt.Println("Finding Weather...")
	fmt.Print("Which location would you like to find the weather for?")
	reader := bufio.NewReader(os.Stdin)
	var geoCodingEndpoint = "https://geocoding-api.open-meteo.com/v1/search?name=%L"


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

