// constants declared like vars but with keyword const
// can be chars, strings, bookls, numeric
// cant be declared with := syntax

package main

import "fmt"

const pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
