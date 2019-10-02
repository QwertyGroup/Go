package main

import (
	"fmt"
	"time"
)

func main() {
	Run()
}

func deamon() {
	input, output := make(chan int, 10), make(chan int)
	go func() {
		for val := range input {
			output <- val
		}
	}()
	for i := 0; i < 10; i++ {
		input <- i
	}
	for i := 0; i < 10; i++ {
		v := <-output
		fmt.Println(v)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Work - unit of work
type Work struct {
	x, y, z int
}

func (w *Work) String() string {
	return fmt.Sprintf("(x:%v; y:%v; z:%v)", w.x, w.y, w.z)
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(time.Millisecond * time.Duration(w.z))
		out <- w
	}
}

// NumWorkers number of workers
var (
	NumWorkers = 10
)

// Run - run work
func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)

	}
	go sendWork(in)
	receiveWork(out)
}

func sendWork(in chan<- *Work) {
	for i := 0; i < NumWorkers; i++ {
		in <- &Work{x: 1, y: 2}
	}
}

func receiveWork(out <-chan *Work) {
	for i := 0; i < NumWorkers; i++ {
		w := <-out
		fmt.Println(w)
	}
}
