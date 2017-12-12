// when slicing, omit high or low bounds to use defaults

// default is zero for low, and length for high

//For array:
// var a [10]int

//These are all the same:
//a[0:10]
//a[:10]
//a[0:]
//a[:]

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
