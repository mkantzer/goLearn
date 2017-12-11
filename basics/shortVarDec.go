// inside a func, the := short statement
// replaves var dec with implicit type
//outside, every statement begins with a keyword, so not available

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
