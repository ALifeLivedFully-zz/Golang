package main

import ("fmt")

const usixteenbitmax float64 = 65535
const kmhmultiplier float64 = 1.60934

type car struct {
	gaspedal,breakpedal uint16 //min 0 max 65,535
	steeringwheel int16
	topspeedkmh float64
}

func (c car) kmh() float64 {
	return float64(c.gaspedal) * (c.topspeedkmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gaspedal) * (c.topspeedkmh/usixteenbitmax/kmhmultiplier)
}

func (c *car) newtopspeed (newspeed float64) {
	c.topspeedkmh = newspeed
}

func main()  {
	car1 := car{65000,0,12561,225.0}
	fmt.Println("all car data: ",car1)
	fmt.Println("gas pedal data: ",car1.gaspedal)
	fmt.Println("break pedal data: ",car1.breakpedal)
	fmt.Println("steering data: ",car1.steeringwheel)
	fmt.Println("top speed data: ",car1.topspeedkmh)
	fmt.Println(car1.kmh())
	fmt.Println(car1.mph())
	car1.newtopspeed(500)
	fmt.Println(car1.kmh())
	fmt.Println(car1.mph())
}