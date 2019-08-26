// goroutines
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
}

func one() {
	go say("Hey") // lightweight thread
	say("There")  // main thread locker
}

func two() {
	go say("Hey")           // lightweight thread
	go say("There")         // lightweight thread
	time.Sleep(time.Second) // main thread locker
}

func main() {
	// one()
	two()
}
