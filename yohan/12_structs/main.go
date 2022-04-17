package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastname
	}
}

func main() {
	// init person using struct

	person1 := Person{firstName: "Yohan", lastName: "Sun", city: "Boston", gender: "m", age: 30}
	person2 := Person{"Hanna", "Yoon", "Boston", "f", 30}

	fmt.Println(person2)
	fmt.Println(person2.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday() // 1++
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(person2)
	person2.getMarried("Kim")
	fmt.Println(person2)
	fmt.Println(person2.greet())

}
