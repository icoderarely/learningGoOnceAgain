package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i // shadows the outer scope total variable
		fmt.Println(total)
	}
	fmt.Println(total) // it prints 0, as it is the int' zero value
}
