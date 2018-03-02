// no classes, but you can define methods on type
//methods are functions with a /receiver/ arg

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//receiver appears betewen func and method name: here, v Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Abs method is a receiver of type Vertex, named v

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// to write it as a regular function:
//func Abs(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//rememeber, method just a function with a receiver arg
