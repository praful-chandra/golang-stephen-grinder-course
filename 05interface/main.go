package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	side int
}
type square struct {
	side int
}

func (t triangle) getArea() float64 {
	return 0.5 * (float64(t.side) * float64(t.side))
}

func (s square) getArea() float64 {
	return float64(s.side * s.side)
}

func printArea(s shape) {
	fmt.Println("Area of  is ", s.getArea())
}

func main() {

	tri := triangle{side: 10}
	sqr := square{side: 10}

	printArea(tri)
	printArea(sqr)

}
