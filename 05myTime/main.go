package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time!")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println("This is created time!")
	createdDate := time.Date(2035, time.September, 20, 20, 20, 0, 0, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
