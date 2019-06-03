package main

import (
	"fmt"
)

func main() {
	x := 15 // sets variable x equal to 15
	a := &x // sets variable a equal to the memory address of x 

	fmt.Println(x) // prints variable x which is equal to 15
	fmt.Println(*a) // * reads thru variable a which leads to variable x which is 15
	fmt.Println(a) // prints variable a which is equal to the memory address of variable x
}