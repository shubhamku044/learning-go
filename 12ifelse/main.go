package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in Go!")

	loginCount := 25

	if loginCount < 10 {
		fmt.Println("Login count is less than 10")
	} else if loginCount < 20 {
		fmt.Println("Login count is less than 20")
	} else {
		fmt.Println("Login count is greater than 20")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}
