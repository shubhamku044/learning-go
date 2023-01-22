package main

import "fmt"

func main() {
	fmt.Println("Welcome to the arrays in go language")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Length of Fruit List: ", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Veg List: ", vegList)
	fmt.Println("Length of Veg List: ", len(vegList))
}
