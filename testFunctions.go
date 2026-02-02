package main

import("fmt")

func main(){
	var name string = "Unedited Name"

	func helper(){
		name = "Edited Name"
	}

	helper()
	fmt.Println(name)
}
