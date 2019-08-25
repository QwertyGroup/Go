// methods: pointer receivers
package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMult float64 = 1.60934

type car struct {
	gasPedal      uint16 // min 0 max 65535
	brakePedal    uint16
	steeringWheel int16 // -32k to 32k
	topSpeedKmh   float64
}

func (c car) kmh() float64 { // value receiver method
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	c.topSpeedKmh = 123
	return c.kmh() / kmhMult
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	aCar := car{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0}
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
