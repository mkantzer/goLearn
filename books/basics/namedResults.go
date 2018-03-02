// can name return values, defined at top of function
// document meaning

// return statement without args returns named return values
// this is really bad form outside of really small functions

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(17))
}
