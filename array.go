package main

import "fmt"

func main() {
	var a [3]string
	a[2] = "Hello"
	a[1] = "World"
	a[0] = "!"
	fmt.Println(a[2], a[1], a[0])
	fmt.Println(a)

	primes := [7]int{2, 3, 5, 7, 11, 13} // 2,3,5,7,11,13,0
	fmt.Println(primes)
}
