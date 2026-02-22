
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

				//Concurrent Test
				fmt.Println("---- Starting Concurrent Test ----")
				concurrentStart := time.Now()
				for _, city := range(cities){
					city = strings.TrimSpace(city)
					wg.Add(1)
					go func(c string){
						defer wg.Done()
						weatherlib.FindWeather(c)
					}(city)
				}
				wg.Wait()
				fmt.Printf("Finding Weather concurrently took %v \n", time.Since(concurrentStart))

				//Sequential/Single Threaded Test
				fmt.Println("\n ---- Starting Sequential Test ----")
				sequentialStart := time.Now()
				for _, city := range(cities){
					city = strings.TrimSpace(city)
					weatherlib.FindWeather(city)
				}
				fmt.Printf("Finding Weather sequentially took %v \n", time.Since(sequentialStart))
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

