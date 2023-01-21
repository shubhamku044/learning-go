package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "This is a simple welcome message"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	// comma ok || error ok
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered:", text)
	fmt.Printf("Variable text of type: %T \n", text)

}
