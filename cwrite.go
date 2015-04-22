package main

import (
	"fmt"
	"sync"
)

var testString = "likewtfomgbbq"
var steps = 100000
var sampleSize = 10

func TestRun(store Store) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < steps; i++ {
			store.Put(testString)
		}
		wg.Done()
	}()

	go func() {
		misses := 0
		for i := 0; i < steps; i++ {
			current := store.Get()
			if current != testString {
				misses++
			}
		}
		//fmt.Printf("Misses: %d", misses)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	fmt.Printf("Going for it\n")

	store := NewSwapStore(sampleSize)
	//store := NewChanneledStore(sampleSize)
	TestRun(store)
}
