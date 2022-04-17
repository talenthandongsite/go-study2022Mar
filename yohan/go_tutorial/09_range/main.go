package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 5}

	// Loop through ids
	for ik123, idk123 := range ids {
		fmt.Printf("%d - ID: %d\n", ik123, idk123)
	}

	// Not using index
	for _, idk123 := range ids {
		fmt.Printf("ID: %d\n", idk123)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
