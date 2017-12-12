// slices are like refernces to arrays

// dont store any data, just describes section of an array

//change the slice, change the array.
/// other slices of same array will see those

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}
