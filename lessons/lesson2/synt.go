package main

import (
	"fmt"
	"math"
	"math/rand"
)

func showRoot() {
	fmt.Println("The square root of 5 is ", math.Sqrt(5))

}

func showRnd() {
	fmt.Println("Rnd:", rand.Intn(100))
}

func main() {
	showRoot()
	showRnd()
}
