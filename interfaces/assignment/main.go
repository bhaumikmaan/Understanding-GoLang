package main

import "fmt"

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{side: 2}
	tr := triangle{height: 3, base: 5}
	print(sq)
	print(tr)
}

func print(sh shape) {
	fmt.Println("Area for the shape ", sh, " is ", sh.getArea())
}

func (sq square) getArea() float64 {
	return sq.side * sq.side
}

func (tr triangle) getArea() float64 {
	return tr.height * tr.base * 0.5
}
