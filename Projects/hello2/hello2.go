package main

import (
	"fmt"
)

func greet(name string) string {
	return "Hello " + name
}

func main()  {
	var name = "first"
	fmt.Println("please enter your name")
	fmt.Scanln(&name)
	fmt.Println(greet(name))
}