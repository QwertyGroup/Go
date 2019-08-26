package main

import (
	"fmt"
)

func zuu(c chan int, val int) {
	c <- val * 5
}

func main() {
	zuuchan := make(chan int)
	go zuu(zuuchan, 3)
	go zuu(zuuchan, 5)
	go zuu(zuuchan, 1<<16)
	v1, v2, v3 := <-zuuchan, <-zuuchan, <-zuuchan // This has to happend. Blocking
	fmt.Println(v1, v2, v3)
}
