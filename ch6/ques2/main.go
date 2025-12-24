package main

import "fmt"

// Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position in the passed-in slice to the passed-in string. At the end of UpdateSlice, print the slice after making the change. The GrowSlice function also takes in a []string and a string. It appends the string onto the slice. At the end of GrowSlice, print the slice after making the change. Call these functions from main. Print out the slice before each function is called and after each function is called. Do you understand why some changes are visible in main and why some changes are not?

func UpdateSlice(slc []string, str string) {
	slc[len(slc)-1] = str
	fmt.Println("inside UpdateSlice:", slc)
}

func GrowSlice(slc []string, str string) {
	slc = append(slc, str)
	fmt.Println("inside GrowSlice:", slc)
}

func main() {
	slc := []string{"a", "b", "c"}

	fmt.Println("before UpdateSlice:", slc)
	UpdateSlice(slc, "d")
	fmt.Println("after UpdateSlice:", slc)

	fmt.Println("before GrowSlice:", slc)
	GrowSlice(slc, "e")
	fmt.Println("after GrowSlice:", slc)
}
