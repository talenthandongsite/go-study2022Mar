package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	carArr := [2]string{"Hyundai", "BMW"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	fmt.Println(carArr)
	fmt.Println(carArr[0])
	fmt.Println(carArr[1])

	fruitSlice := []string{"Apple", "Orange", "Cherry"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
	fmt.Println(fruitSlice[1:3])
}
