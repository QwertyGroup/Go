package main

import "fmt"

func main() {
	res := gen(3)
	fmt.Println(res)
	fmt.Println(calc(res))
}

func gen(n int) *[]float64 {
	c := make([]float64, n)
	i := 0
	for sum(&c) >= prod(&c) {
		c[i] += 0.01
		i++
		i %= len(c)
	}
	return &c
}

func sum(s *[]float64) (acc float64) {
	for _, v := range *s {
		acc += v
	}
	return
}

func prod(s *[]float64) float64 {
	acc := 1.0
	for _, v := range *s {
		acc *= v
	}
	return acc
}

func calc(s *[]float64) (res float64) {
	for _, v := range *s {
		res += 1 / v
	}
	return
}
