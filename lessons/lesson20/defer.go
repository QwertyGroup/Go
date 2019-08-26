package main

import (
	"fmt"
	"sync"
	"time"
)

func goo() {
	defer fmt.Println("Done!")        // defer run on success or on panic
	defer fmt.Println("Are we done?") // LIFO
	fmt.Println("Doing some stuff")
}

func boo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done() // Place it here
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
	// wg.Done() // Remove from here and
}

func zuu() {
	wg.Add(2)
	go say("Go")
	go say("Routine")
	wg.Wait()
}

func main() {
	// goo()
	// boo()
	zuu()
}
