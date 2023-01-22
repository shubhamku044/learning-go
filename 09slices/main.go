package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go!")

	var fruitList = []string{"Apple", "Orange", "Grape", "Plum"}

	fmt.Println("Fruit List:", fruitList)
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	// append
	fruitList = append(fruitList[1:])
	fmt.Println("Fruit List:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 9999
	highScores[1] = 9998
	highScores[2] = 9997
	highScores[3] = 9996

	fmt.Println("High Scores:", highScores)

	// append
	highScores = append(highScores, 9995, 9994, 9993, 9992, 9991, 9990)
	fmt.Println("High Scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("High Scores:", highScores)

	fmt.Println("---------------------")
	// remove values from slice
	var courses = []string{"Docker", "Kubernetes", "Puppet", "Terraform", "AWS", "Azure", "GCP", "Python", "Go"}
	fmt.Println("Courses:", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses:", courses)
	fmt.Println("Length of courses:", len(courses))

}
