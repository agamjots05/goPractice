package main

import("fmt")

func main(){
	ch := make(chan string)	

	go helper(ch)

	fmt.Println("Waiting for channel to recieve data")
	ch <- "Sent data now"

	fmt.Println("Wait 1")
	fmt.Println("Wait 2")


	fmt.Println("Wait 3")




}

func helper(channel chan string){
	fmt.Println(<- channel)

}
