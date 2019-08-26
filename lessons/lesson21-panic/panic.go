package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanp:", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func zuu() {
	wg.Add(2)
	go say("Go")
	go say("Routine")
	wg.Wait()
}

func main() {
	zuu()
}
