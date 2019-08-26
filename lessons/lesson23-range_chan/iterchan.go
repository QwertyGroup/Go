package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func zuu(c chan int, val int) {
	defer wg.Done()
	c <- val * 5
}

func main() {
	count := 10
	zuuchan := make(chan int, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go zuu(zuuchan, i)
	}

	wg.Wait()
	close(zuuchan)

	for item := range zuuchan {
		fmt.Println(item)
	}
}
