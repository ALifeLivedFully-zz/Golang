package main

import (
	"math"
	"fmt"
)

func full(z, w string) string {
	return z+w
}

func add(x, y float64) float64 {
	return x+y
}

func main() {
	var num1,num2,num3 float64
	var name1,name2 string
	fmt.Printf("first name?")
	fmt.Scanln(&name1)
	fmt.Printf("last name?")
	fmt.Scanln(&name2)
	fmt.Printf("ok %s %s, first number? ",name1,name2)
	fmt.Scanln(&num1)
	fmt.Printf("second number?")
	fmt.Scanln(&num2)
	fmt.Println("Calculating...")
	fmt.Println(full(name1, name2))
	fmt.Println(add(num1,num2))
	fmt.Println("choose a number to square:")
	fmt.Scanln(&num3)
	fmt.Println("Calculating...")
	fmt.Println(math.Sqrt(num3))
}
