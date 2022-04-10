package main

import "fmt"

func main() {
	// var name = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3
	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"
	name, email := "Brad", "brad@gmail.com"
	// size := 1.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

}
