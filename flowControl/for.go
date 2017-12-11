// Go has only one loop: for
// the basic for loop, has 3 semi-colon statements:
// init: executed before first iteration
// condition: evaluated before every iteration
// post: executed at the end of every iteration

// init is usually short var declaration. vars from this are only in scope of for

// loop stops when condition evaluates as false

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println("Final = ", sum)
}
