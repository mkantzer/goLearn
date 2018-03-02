// https://tour.golang.org/moretypes/18

// Implement Pic:
// Return: Slice of length 'dy', each element of which is:
// slice of dx 8-bit unsigned

// when you run the progrma, it will display your picture, interpreting ints as grayzscale calues

//choice of image is up to you but (x+y)/2, x*y, and x^y are interesting

//(You need a loop to allocate each []uint8 inside [][]uint8)

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
