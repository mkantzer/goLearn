// https://tour.golang.org/moretypes/2
// struct is a collection of fields

// fields accessed using a dot

// or, they can be accessed through a struct pointer:
// To access X of a struct when have struct point P, we could do:
// (*p).X , but that suckssss

//instead, we can just write p.X

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
