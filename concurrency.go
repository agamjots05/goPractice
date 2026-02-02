package main

import(
	"fmt"
	"time"
)

func helper(msg string){
	fmt.Println(msg)

	time.Sleep(time.Second * 1)

	fmt.Println("1 Second Has passed")
}

func main(){
	go helper("Method Call 1")
	helper("Method Call 2")

	
	chn1 := make(chan string)
	fmt.Println(<-chn1)


}
