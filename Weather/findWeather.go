package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"net/http"
	"net/url"
)

func findWeather(){
	fmt.Print("Which location would you like to find the weather for? ")
	reader := bufio.NewReader(os.Stdin)


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

		// This will transform the spaces into "%20" e.g. "san jose" --> "san%20jose"
		escapedLocation := url.QueryEscape(location)

		urlLocation := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", escapedLocation)
		resp, err := http.Get(urlLocation)
		if err != nil{
			fmt.Println("Error While Requesting Location. \n Please Try Again")
		}
		fmt.Println("Response: ", resp)
		
			
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

