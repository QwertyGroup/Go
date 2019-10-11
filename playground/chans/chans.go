package main

import (
	"fmt"
	"sync"
	"time"
)

// Error is error level log
// Warning is warning level log
// Info is info level log
const (
	Error   logLevel = "ERR"
	Warning logLevel = "WRN"
	Info    logLevel = "INF"
)

type logLevel string

type logEntry struct {
	time    time.Time
	level   logLevel
	message string
}

var wg = sync.WaitGroup{}
var logCh = make(chan logEntry)
var doneCh = make(chan struct{})

func main() {
	go logger(logCh, doneCh)
	defer func() { doneCh <- struct{}{} }()

	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // receive-only (from) chan (read)
		for val := range ch {
			logCh <- newInfoEntry(fmt.Sprintf("Read %v from chan", val))
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // send-only (to) chan (write)
		val := 42
		logCh <- newInfoEntry(fmt.Sprintf("Send %v to chan", val))
		ch <- val
		val = 21
		logCh <- newInfoEntry(fmt.Sprintf("Send %v to chan", val))
		ch <- val
		close(ch)
		logCh <- newInfoEntry("Closed chan")
		wg.Done()
	}(ch)
	logCh <- newInfoEntry("Waiting goroutines...")
	wg.Wait()
}

func newInfoEntry(msg string) logEntry {
	return logEntry{time: time.Now(), level: Info, message: msg}
}

func logger(logCh <-chan logEntry, doneCh <-chan struct{}) {
	log := func(entry logEntry) {
		fmt.Printf("[%v %v] :: %v\n",
			entry.time.Format("15:04:05"),
			entry.level, entry.message)
	}
	for {
		select {
		case entry := <-logCh:
			log(entry)
		case <-doneCh:
			log(logEntry{time: time.Now(), level: Warning, message: "logging finished"})
			break
		}
	}
}
