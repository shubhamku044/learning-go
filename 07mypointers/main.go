package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers class")

	// var ptr *int
	// fmt.Println("The default value of ptr is", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The value of ptr is", ptr)
	fmt.Println("The value of myNumber is", myNumber)
	fmt.Println("The value of *ptr is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("The new value of *ptr is", *ptr)
	fmt.Println("The new value of myNumber is", myNumber)
}
