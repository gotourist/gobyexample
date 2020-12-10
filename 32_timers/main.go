package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	time.Sleep(2 * time.Second)
	stop2 := timer2.Stop()
	//var stop2 bool
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}
