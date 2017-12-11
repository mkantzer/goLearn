// vars declared but not given init are given the zero value
// 0 for numbers
// false for bools
// "" (empty string) for strings

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
