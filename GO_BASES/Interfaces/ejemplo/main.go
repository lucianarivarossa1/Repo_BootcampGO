package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() (r float64)
	perim() (r float64)
}
type rectangulo struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangulo) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func details(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())

}
func main() {
	//r := rectangulo{width: 2, height: 5}
	//c := circle{radius: 5}
}
