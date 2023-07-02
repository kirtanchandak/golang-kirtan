package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers!!")

	mynum := 25

	var ptr = &mynum

	fmt.Println("The adderess of myNum is: ", ptr) //0xc000016098
	fmt.Println("The value of myNum is: ", *ptr)   // 25

	//Notes -
	// & - is used to refrence to a variable
	// * - is used to point the value of pointer.
}
