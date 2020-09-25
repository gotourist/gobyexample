package main

import "fmt"

func fact(a int) int {
	if a == 1 {
		return 1
	}
	return a * fact(a-1)
}

func main() {
	var n int
	fmt.Scanln(&n)

	fmt.Println(fact(n))

}
