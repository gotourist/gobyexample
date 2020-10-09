package main

import (
	"fmt"
	"time"
)

func funny(from string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(from, ":", i)

	}
}
func main() {
	go funny("Hello")
	go funny("goroutine")
	funny("direct")

	time.Sleep(time.Second)
	fmt.Println("done")
}
