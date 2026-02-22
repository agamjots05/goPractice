package weatherlib

import (
	"fmt"
	"net/http"
	"net/url"
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
type WeatherResponse struct {
	Results struct {
		Temperature float32 `json:"temperature_2m"`
	} `json:"current"`
	

}

func FindWeather(location string){
		// This will transform the spaces into "%20" e.g. "san jose" --> "san%20jose"
		escapedLocation := url.QueryEscape(location)

		urlLocation := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", escapedLocation)
		resp, err := http.Get(urlLocation)
		fmt.Println("Finding Latitude and Longitude...")
		if err != nil{
			fmt.Println("Error While Requesting ", location)
			return
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()


		// Convert the body from RAW JSON to a readable format using the 'GeoRespose' we created
		var geo GeoResponse
		err = json.Unmarshal(body, &geo)
		if err != nil {
			fmt.Println("Error Unmarshalling ", location)
			return
		}
		if len(geo.Results) == 0{
			fmt.Println("Error finding coordinates for ", location)
			return
		}
		lat := geo.Results[0].Latitude
		lon := geo.Results[0].Longitude


		//Use lat and lon we found to find the final temp for location
		urlTemperature := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m&current=temperature_2m&timeformat=unixtime&temperature_unit=fahrenheit", lat, lon)
		resp, err = http.Get(urlTemperature)
		if err != nil {
			fmt.Println("Error Getting Temperature for ", location)
			return
		}
		body, _ = io.ReadAll(resp.Body)
		resp.Body.Close()

		var temp WeatherResponse
		err = json.Unmarshal(body, &temp)
		if err != nil {
			fmt.Println("Error unmarshalling final temperature for ", location)
		}
		finalTemp := temp.Results.Temperature
		fmt.Printf("Final Temperature: %f Degrees Farenheit For Location: %s \n", finalTemp, location)
		return 

}
