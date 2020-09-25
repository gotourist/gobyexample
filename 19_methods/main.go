package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height int
}

type cir struct {
	radius float64
}

func (r *rect) area() int {
	return r.width * r.height
}

func (c *cir) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	c := cir{radius: 2}
	fmt.Println("area: ", r.area())
	fmt.Println("area: ", c.area())

	// fmt.Println("perim:", r.perim())

	// rp := &r
	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim:", rp.perim())
}
