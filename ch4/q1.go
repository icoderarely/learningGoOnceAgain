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
	fmt.Println(sl)
}
