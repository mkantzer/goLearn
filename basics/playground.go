package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground")
	for i := 0; i <= 10; i++ {
		fmt.Println("the time is", time.Now())
		time.Sleep(100)
	}
}
