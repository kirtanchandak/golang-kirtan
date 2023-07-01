package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome!")
	fmt.Println("Please rate us between 1 and 5 ⭐")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("You rated us", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating", numRating+1)
	}
}

// Notes -
// strconv - is used to convert stings to int, float etc.

// strings - in this example we have used strings to trim the space from in the input.
// for eg. 5\n got converted to 5
