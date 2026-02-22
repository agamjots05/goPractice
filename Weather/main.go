
package main

import(
	"weather/weatherlib"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"sync"
)
func main(){
	var userNumChoice int
	fmt.Println("Welcome to the Multi-Weather Terminal \n What would you like to do")
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
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Enter Cities (e.g., London, San Jose, New York")
			var wg sync.WaitGroup
			if scanner.Scan(){
				input := scanner.Text()
				cities := strings.Split(input, ",")

				for _, city := range(cities){
					city = strings.TrimSpace(city)
					wg.Add(1)
					start := time.Now()
					go func(curCity string){
						defer wg.Done()
						weatherlib.FindWeather(city)
						fmt.Printf("Finding Weather for %s took %v \n", city, time.Since(start))
					}(city)
				}
				wg.Wait()

			}
			return
		case 2:
//			learnNewAreas()
			return
		default:
			fmt.Println("Invalid Number. Try again")
		}
	}
}

