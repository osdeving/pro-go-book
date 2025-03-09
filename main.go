package main

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mux = sync.Mutex{}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		mux.Lock()
		*val++
		mux.Unlock()
	}
	wg.Done()
}

func main() {
	counter := 0

	numRoutines := 3
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	wg.Wait()

	Printfln("Total: %v", counter)
}
