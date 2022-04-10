package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string) // [key] value

	// Assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails) // map[Bob:bob@gmail.com Sharon:sharon@gmail.com]
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"]) // bob@gmail.com

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
