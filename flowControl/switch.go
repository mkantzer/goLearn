//switch is short way to write if / else
//runs first case where val = condition

//Like C,C++,Java, JS, PHP, but:
// Go only runs selected, not all after
//break is provided automatically

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// free/open bsd, plan9, windows....
		fmt.Printf("%s.", os)
	}
}

// switch is evaluated from top to bottom, stopping when one succeeds

// switch with no condition is run as switch true
