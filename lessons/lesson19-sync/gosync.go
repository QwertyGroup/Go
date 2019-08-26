// goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go say("Hey")
	go say("There")
	wg.Wait()
	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
