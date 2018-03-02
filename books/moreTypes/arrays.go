//Arrays:

// Type [n]T is an array of n values of type T
// Expression:
// var a [10]int
// declares variable a as array of 10 integers

//array length is part of type, so NO resizing.
// "Dont worry, Go provides convenient way to work with arrays"

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
}
