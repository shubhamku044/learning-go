package main

import "fmt"

// noOfUsers := 100  -- this is not allowed
const LoginToken string = "1234" // this is Public Constant

func main() {
	var username string = "Shubham"
	fmt.Println(username)
	fmt.Printf("Variable username of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable isLoggedIn of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable smallVal of type: %T \n", smallVal)

	var smallFloat float32 = 255.25545325234532
	fmt.Println(smallFloat)
	fmt.Printf("Variable smallFloat of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable anotherVariable of type: %T \n", anotherVariable)

	var anotherFloat float64
	fmt.Println(anotherFloat)
	fmt.Printf("Variable anotherFloat of type: %T \n", anotherFloat)

	var anotherBool bool
	fmt.Println(anotherBool)
	fmt.Printf("Variable anotherBool of type: %T \n", anotherBool)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable anotherString of type: %T \n", anotherString)

	// implicit type declaration
	var implicitString = "Implicit String"
	fmt.Println(implicitString)
	fmt.Printf("Variable implicitString of type: %T \n", implicitString)
	// impliedString = 2 -- this is not allowed

	// short hand declaration
	shortHandString := "Short Hand String"
	fmt.Println(shortHandString)
	fmt.Printf("Variable shortHandString of type: %T \n", shortHandString)

	// short hand declaration
	numOfUsers := 100
	fmt.Println(numOfUsers)

	fmt.Println(LoginToken)
}
