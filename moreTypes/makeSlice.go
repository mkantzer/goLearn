// slices can be created with build-in `make`
// how you create dynamically sized arrays

// make function allocates a zeroed array, returns a slice that refers to it:

//a := make([]int, 5) // len(a) = 5

//To specify a capacity, pass a third argument:
// b:= make([]int, 0, 5) //len(b)=0, cap(b)=5

//b = b[:cap(b)]  	// len(b)=5, cap(b)=5
//b = b[1:] 		// len(b)=4, cap(b)=4

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("D", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
