package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("You entered:", input)

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Error while parsing input", err)
	} else {
		fmt.Println("Your rating is:", numRating+1)
	}
}
