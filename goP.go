package main
import("fmt")

func main(){
	// Testing print statements
	fmt.Println("test")	

	// Testing setting variables
	// Two ways to initialize variable 
	// Hard typing the variable manually
	/* 
		1. There aren't any restrictions when it comes to type casting var's manually 
		2. Manual type casting allows you to not assign a value to a var after it's been declared
			eg. var coolName string
				coolName = "ReallyCoolName"

	*/
	var name1 string = "Alex"

	//Letting compiler decide the type
	/* 
		There are some rules when it comes to this kind of declaration
			1. This can only be used within functions
			2. You MUST assign a value after it's been declared
	*/
	name2 := "Susan"


	fmt.Println(name1)
	fmt.Println(name2)

	//Constants 
	fmt.Println("Constant's Sections")

	//Constants are declared with a 'const' before the constant's name 
	const pi float32 = 3.14
	
	//Constants can also written without type casting them
	const e = 2.71

	fmt.Printf("The value of pi = %v and the type is %T \n",pi,pi)
	fmt.Printf("The value of e = %v. Type of e is %T \n",e,e)

	
	//Data Type Section
	fmt.Println("Data Type Section")
	// Go has 3 major data types being 
	// Numerical 
//	var num1 int = 123
//	var num2 float32 = 12.2

	//String
//	var str1 string = "stringOne"

	//Boolean
//	var isTrue bool = false


	//Arrays 

	fmt.Println("Arrays and Slices Section")
	// Arrays are defined by this formula var arr = [length Of Array]dataType{hard coded elements}
	//Initialize an array where you can define length
	var arr = []int{1,2}
	fmt.Printf("arr = %v", arr)

	arr = append(arr, 3)
	//fmt.Printf(arr)
	//Initialize an array where the length is determined by compiler
	var arr2 = [...]int{1,3,5,7}
	fmt.Printf("arr2 = %v",arr2)

	/*	
		Slices are very similiar to arrays in the sense that they act as ArrayLists in java. They are mutable in size
		You define slices exactly like arrays, however you leave the length of the slice empty to signify the instantiation of a slice

		If you create a slice through an array undertand that the slice is just a pointer to the original array
			eg slicing a pre-ex array --> var arr = [2]int{1,2}; var slice = arr[0:1]
			In this case the slice == [1] since the starting index is inclusive and the ending is non-inclusive
		
		The capacity of this slice will be 2 and the length is 1G

	*/	

	var arr3 = []int{1,2,3,4,5}
	fmt.Printf("arr3 = %v \n", arr3)
	
	//Conditional Statements
	fmt.Println("Conditional Statements Section")

	//If Statements
	isCool := true
	isSmart := false
	if isCool && isSmart{
		fmt.Println("Person is cool and smart")
	}	else{
		fmt.Println(":(")
	}



	//For loops

	fmt.Println("For Loop section")
	for i:= 0; i < 5; i++{
		fmt.Println(i)
	}
	

	//Go Pointers Section
	fmt.Println("Pointer Section")
	var actualValue = "Cool Value"
	fmt.Println("Memory Address of our variable: ", &actualValue)


	//You can also create variables that store memory addresses
	var pointerVar *string = &actualValue

	//You can get the VALUE of the pointer by adding * in front of the pointer
	fmt.Println("Value of Pointer : ", *pointerVar)
	fmt.Println("Memory Address of Pointer:", &pointerVar)

	var number int = 10
	var pntrNum *int = &number

	//Since the the pointer is pointing at the var 'number' You can change the value of 'number' using 'pntrNum'
	fmt.Println("Original Value of number", number)
	*pntrNum = 15
	fmt.Println("New Value for pntrNum after changing number from 10 to 15", number)

	
	//Go routines section
	//The idea of a go routine is to allow cocurreny using Go



}


