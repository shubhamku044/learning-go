package main

import "fmt"

func main() {
	fmt.Println("Welcome to for loop in Go!")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("------------")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("------------")
	for index, day := range days {
		fmt.Printf("Index: %d, Day: %s\n", index, day)
	}

	fmt.Println("------------")

	rougeValue := 0
	for rougeValue < 10 {
		fmt.Println(rougeValue)
		if rougeValue == 3 {
			goto lco
		}
		if rougeValue == 5 {
			break
		}
		rougeValue++
	}

	// Goto
	fmt.Println("------------")
lco:
	fmt.Println("Welcome to LCO")
}
