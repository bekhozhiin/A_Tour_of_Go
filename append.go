package main

import "fmt"

func main() {
	var s []int
	printSlice(s) // len and cap=0

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) // len and cap=1

	// The slice grows as needed.
	s = append(s, 1, 6, 5)
	printSlice(s) //len=4 and cap=4

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s) // len=7 and cap=8 (because, when len>cap--> cap*2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
