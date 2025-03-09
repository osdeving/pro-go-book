package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}

var squares = map[int]int{}

var src = rand.NewSource(time.Now().UnixNano())
var rnd = rand.New(src)

func calculateSquares(max, interations int) {
	for i := 0; i < interations; i++ {
		val := rnd.Intn(max)

		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()

		if ok {
			fmt.Printf("Cached value: %v = %v\n", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				fmt.Printf("Added value: %v = %v\n", val, squares[val])
			}
			rwmutex.Unlock()
		}

	}
	wg.Done()
}

func main() {
	numRoutines := 3

	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go calculateSquares(10, 5)
	}
	wg.Wait()
	fmt.Printf("Cached values: %v\n", len(squares))
}
