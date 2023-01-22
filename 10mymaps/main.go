package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Go!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Go"
	languages["RB"] = "Ruby"
	languages["PHP"] = "PHP"

	fmt.Println("Languages:", languages)

	// get value
	fmt.Println("JS:", languages["JS"])

	// delete
	delete(languages, "RB")
	fmt.Println("Languages:", languages)

	// loops on maps
	for key, value := range languages {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
