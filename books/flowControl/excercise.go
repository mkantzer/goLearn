// let's implement a sqrt function:
// given number x, find z for which z^2 most close to x

//usually, done with a loop
//start with some guess z, can adjust based on how close
// z^2 is to x, producing better guess

// z -= (z*z - x) / (2*z)

// repeat this, to make guess better until reach as close as can be

// implement this in func Sqrt provided.

// good start is z = 1
// repeat 10 times and print each z along the way,
//see how close you can get to various x
//

package main

import (
	"fmt"
	"math"
)

func banana(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		n := z - ((z*z - x) / (2 * z))
		dif := math.Abs(z - n)
		if dif < .0000000000001 {
			return z
		}
	}

	return z
}

func main() {

	const x float64 = 157.8
	fmt.Println(
		banana(x),
		math.Sqrt(x),
		banana(x)-math.Sqrt(x),
	)

}
