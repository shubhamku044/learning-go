package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in go lang")

	presentTime := time.Now()

	fmt.Println("Present time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.September, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date is ", createdDate)
	fmt.Println("Created date is ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
