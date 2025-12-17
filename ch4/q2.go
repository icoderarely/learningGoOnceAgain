package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sl := make([]int, 100)
	for i := range 100 {
		sl[i] = rand.Int()
	}
	for _, v := range sl {
		if v%2 == 0 {
			fmt.Println("Two!")
		} else if v%3 == 0 {
			fmt.Println("Three")
		} else if v%6 == 0 {
			fmt.Println("Six!")
		} else {
			fmt.Println("Never mind!")
		}
	}
}
