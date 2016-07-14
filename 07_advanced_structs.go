package main

import "fmt"

type Common interface {
	setArea()
}

type Square struct {
	Side int
	Area int
}

func (s Square) District() int {
	return s.Side * 4
}

func (s *Square) setArea() {
	s.Area = s.Side * s.Side
}

type Rectangle struct {
	SideA int
	SideB int
	Area int
}

func (r *Rectangle) setArea() {
	r.Area = r.SideA * r.SideB
}

func main()  {
	sqr := Square{Side: 5}
	fmt.Println(sqr.District())

	calculateArea(&sqr)
	fmt.Println(sqr.Area)

	rectangle := &Rectangle{SideA: 5, SideB: 20}
	calculateArea(rectangle)
	fmt.Println(rectangle.Area)
}

func calculateArea(object Common) {
	object.setArea()
}