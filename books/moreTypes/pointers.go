//there are pointers. they hold the memory address of a value

//the type *T is a pointer to a T value. zero value is nil

//var p *int

//& operator generates a pointer to its operand

//i := 42
// p - &i

//the * operator denotes underling value

//fmt.Println(*p) //read i through the pointer p
//*p = 21 //set i through the pointer p

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //point to j
	*p = *p / 37   // divide j though the pointer
	fmt.Println(j) //see the new value of j
}
