package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Go!")

	var gymList [4]string
	gymList[0] = "Lat Pull Down"
	gymList[1] = "Hammer Curl"
	gymList[3] = "Leg Press"

	fmt.Println("Excercises: ", gymList)

	var supplements = [3]string{"Whey Protein", "Creatin", "Mass Gainer"}
	fmt.Println("Supplements: ", supplements)
}
