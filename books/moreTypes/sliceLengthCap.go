// slice has both LENGTH, and CAPACITY

//length = number of elements in slice
//capacity = number of elements in underlying array, from first element in slice

// len(s), cap(s)

//can extend a slices length by re-slicing, if capacity is bigger.
//Try changing one slice opp in the example program to extend beyond

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//Slice the slice to give zero length
	s = s[:0]
	printSlice(s)

	//extend length
	s = s[:4]
	printSlice(s)

	//drop values
	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
