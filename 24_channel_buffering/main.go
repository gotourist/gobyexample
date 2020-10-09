package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	a := 1
	b := 2

	fmt.Println(a)
	fmt.Println(b)
}
