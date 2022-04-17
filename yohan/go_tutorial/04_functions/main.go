package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World2")
	fmt.Println(greeting("Yohan"))
	fmt.Println(getSum(3, 4))
}
