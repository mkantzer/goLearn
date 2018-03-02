package main

import (
	"fmt"
	"math"
)

//Circle does things
type Circle struct {
	x, y, r float64
}

//Rectangle does things
type Rectangle struct {
	x1, y1, x2, y2 float64
}

//CircleArea gets area of a circle, when passed a pointer to a circle
func CircleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := Distance(r.x1, r.y1, r.x1, r.y2)
	w := Distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//Distance returns distance between (x1,y1 and x2,y2)
func Distance(x1, y1, x2, y2 float64) float64 {
	a := (x2 - x1) * (x2 - x1)
	b := (y2 - y1) * (y2 - y1)
	return math.Sqrt(a + b)
}

//Shape is anything with an Area
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(r.x1, r.y2)
	fmt.Println(CircleArea(&c))
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(totalArea(&r, &c, &c, &c, &r))
}
