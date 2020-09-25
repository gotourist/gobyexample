package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(12, 24)
	fmt.Println("result:", res)

	res = plusplus(32, 24, 35)
	fmt.Println("result 2:", res)
}
