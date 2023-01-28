package main

import "fmt"

func main() {
	fmt.Println("learning functions")

	greeter()

	result := adder(2, 3)
	fmt.Println(result)

	result = proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}

func greeter() {
	fmt.Println("Namastey from go lang")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	result := 0
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return result
}
