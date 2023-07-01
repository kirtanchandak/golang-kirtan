package main

import "fmt"

const LoginToken string = "djfdj54lj5r4mk4j5" //Public (because starts with capital)

func main() {
	//variables
	var username string = "kirtan"
	fmt.Println(username)
	fmt.Printf("Var is of type: %T \n", username)

	//boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Var is of type: %T \n", isLoggedIn)

	//integrer
	var num int = 200
	fmt.Println(num)
	fmt.Printf("Var is of type: %T \n", num)

	//default values & aliases
	var anothereVariable int
	fmt.Println(anothereVariable) // output: 0
	fmt.Printf("Var is of type: %T \n", anothereVariable)

	var anothereString string
	fmt.Println(anothereString) // output: ""
	fmt.Printf("Var is of type: %T \n", anothereString)

	//implict values
	var food = "pizza"
	fmt.Println(food)

	//valouros opreator
	noOfUsers := 34000
	fmt.Println(noOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Var is of type: %T \n", LoginToken)

}
