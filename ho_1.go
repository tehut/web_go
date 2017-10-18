package main

import (
	"fmt"
)

type square struct {
	length, height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func sharea(sh shape) {
	fmt.Println(sh.area())
}

func (c circle) area() float64 {
	return c.radius * 3.1428
}

func (s square) area() float64 {
	return s.length * s.height
}

func main() {

	sq1 := square{12, 18}
	c1 := circle{9}

	sharea(sq1)
	sharea(c1)
}
