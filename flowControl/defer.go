// a defer defers a func exec until surrounding function returns

// arguments evaluated immediately, but call not exec until suroundding returns

//you can stack them: they get pushed onto a stack.
// executed last-in-first-out
package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
