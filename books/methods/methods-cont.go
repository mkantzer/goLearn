//you can declare method on non-struct types

package main

import (
	"fmt"
	"math"
)

//numeric type MyGloat, with an Abs method
type MyFloat float64

//you can only declare a method with a receiver whose type is defined in same package.
//CANNOT use types defined in another package, inlcuding built-in (like int)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
