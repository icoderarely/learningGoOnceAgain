package main

import "fmt"

func main() {
	var msg string = "Hi 👩 and 👨"
	r := []rune(msg)
	fmt.Printf("%c, %d, %c", r[3], r[3], msg[3])
}
