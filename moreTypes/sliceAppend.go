// https://tour.golang.org/moretypes/15

// appending to slices is common, so its a standard thing

// func append (s []T, vs ...T) []T
// s is a slice of type T, and the rest are T values to append to the slice

//Resulting value is a slice contianing all the elements of the origional slcie plus new values

//if the backing array is too small to fit all the given values, a bigger array will be alocated
// the new slice will point to the newly allocated array

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//append works on nil slices
	s = append(s, 0)
	printSlice(s)

	//slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
