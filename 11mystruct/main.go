package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Go!")

	// no inheritance in Go: No super or parent
	shubham := User{"Shubham", "shubhamku044@gmail.com", true, 20}

	fmt.Println("User:", shubham)
	fmt.Printf("shubham details are: %+v\n", shubham)
	fmt.Printf("Name of user is %v, his email is %v, his age is %v and last is bool value is %v.\n", shubham.Name, shubham.Email, shubham.Age, shubham.Status)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
