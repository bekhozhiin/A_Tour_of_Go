package main

import "fmt"

func main() {
	san := [4]int{
		1,
		3,
		5,
		7,
	}
	fmt.Println(san)

	a := san[0:3] //1 3 5
	b := san[1:3] // 3 5
	fmt.Println(a, b)

	b[1] = 0
	fmt.Println(a, b)
	fmt.Println(san)
}
