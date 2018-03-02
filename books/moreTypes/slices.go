//slices

//arrays have fixed size.
// slices tho, are dynamically sized, flexible view of array contents

// much more common than arrays

//type []T is a slice with elements of type T
// A slice is formed by specifying 2 indicies, hight and low bound by colon
// a[low : high]
// includes first, excludes last
// a[1:4]
// elements 1, 2, 3

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
