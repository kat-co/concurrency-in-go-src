package main

import (
	"sync"
)

func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // <1>
	onceA.Do(initA)                    // <2>
}
