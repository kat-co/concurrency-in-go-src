package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // <1>

	const numGoroutines = 1e4 // <2>
	wg.Add(numGoroutines)
	before := memConsumed() // <3>
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // <4>
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
