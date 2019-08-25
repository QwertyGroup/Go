package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func test() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5
	fmt.Println(add(num1, num2))
}

func add2(x, y float64) float64 {
	return x + y
}

func test2() {
	var num1, num2 = 5.6, 9.5
	fmt.Println(add2(num1, num2))
}

func test3() {
	num1, num2 := 5.6, 9.5
	fmt.Println(add2(num1, num2))
}

const x int = 5

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	test()
	test2()

	a := 12
	b := float64(a)
	fmt.Println(a, b)
}
