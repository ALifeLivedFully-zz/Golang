package main

import (
	"fmt"
)



func main()  {
	x := 10
	y := 15
	//if else conditional below
	if x<=y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	//else if conditional below
	color := "red"

	if color == "red" {
		fmt.Println("color is red")	
	} else if color == "blue" {
		fmt.Println("color is blue")		
	} else {
		fmt.Println("color is not red or blue")
	}
	//switch conditional
	switch color {
	case "red":
		fmt.Println("switch is red")
	case "blue":
		fmt.Println("switch is blue")
	default:
		fmt.Println("switch is not blue or red")
	}
}