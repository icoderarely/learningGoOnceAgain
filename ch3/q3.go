package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	// first instance: using struct literal without names (order matters)
	first := Employee{
		"John",
		"Doe",
		123,
	}

	// second instance: using struct literal with names (order does not matter)
	second := Employee{
		firstName: "Mary",
		lastName:  "Williams",
		id:        111,
	}

	// third instance: using var declaration + populate using dot notation
	var third Employee
	third.firstName = "Charlie"
	third.lastName = "Brown"
	third.id = 321

	// print all three structs
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
