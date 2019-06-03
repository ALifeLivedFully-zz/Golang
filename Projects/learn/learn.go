package main

import "fmt"

const adult float64 = 18
const beer float64 = 21
const military float64 = 16
const rent float64 = 25

type user struct {
	firstname, lastname string
	age float32
}

/*func userinfo()  {
	fmt.Printf("Enter Your First Name: ")
	fmt.Scanln(&firstname)
	fmt.Printf("Enter Your Last Name: ")
	fmt.Scanln(&lastname)
	fmt.Printf("enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Name: %s %s! \n age: %s \n", firstname, lastname, age)
}*/

func main() {
	user1 := user{firstname: "test",lastname:"testing",age:10.0}
	fmt.Printf("Enter Your First Name: ")
	fmt.Scanln(&user1.firstname)
	fmt.Printf("Enter Your Last Name: ")
	fmt.Scanln(&user1.lastname)
	fmt.Printf("enter your age: ")
	fmt.Scanln(&user1.age)
	fmt.Printf("Name: %s %s \nage: %s \n", user1.firstname, user1.lastname, user1.age)
}