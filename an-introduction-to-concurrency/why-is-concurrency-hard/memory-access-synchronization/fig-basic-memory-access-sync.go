package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex // <1>
	var value int
	go func() {
		memoryAccess.Lock() // <2>
		value++
		memoryAccess.Unlock() // <3>
	}()

	memoryAccess.Lock() // <4>
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock() // <5>
}
