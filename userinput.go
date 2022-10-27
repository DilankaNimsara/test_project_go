package main

import (
	"fmt"
	"time"
)

func userInput() {
	var numberOne float64
	var numberTwo float64

	startTime := time.Now()

	fmt.Print("Enter number one : ")
	fmt.Scanln(&numberOne)
	fmt.Println("number one : ", numberOne)
	fmt.Println("Enter number two : ")
	fmt.Scanln(&numberTwo)
	fmt.Println("number two : ", numberTwo)

	fmt.Println("sum : ", numberOne+numberTwo)

	endTime := time.Since(startTime)
	fmt.Println("running time : ", endTime)

}

func findSecondLargest() {
	arr := [10]int{10, 62, 52, 14, 9, 63, 100, 58, 35, 2}
	for i := 0; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

	}

	fmt.Println(arr)
}
