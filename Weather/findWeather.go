package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"net/http"
	"net/url"
	"log"
	"io"
	"encoding/json"
)
type GeoResponse struct {
	Results []struct {
		Name string `json:"name"`
		Latitude float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	}`json:"results"`
}
//type WeatherResponse struct {
//
//}

func findWeather(){
	reader := bufio.NewReader(os.Stdin)
	var location string
	for {
		
		fmt.Print(" Which location would you like to learn about the weather for?")
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
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			log.Fatal(err)
			continue
		}
		// Convert the body from RAW JSON to a readable format using the 'GeoRespose' we created
		var geo GeoResponse
		err = json.Unmarshal(body, &geo)
		if err != nil {
			log.Fatal(err)	
		}
		fmt.Println("Unmarshalled Data", geo)
		break

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

