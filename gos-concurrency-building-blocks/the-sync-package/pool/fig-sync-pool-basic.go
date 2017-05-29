package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // <1>
	instance := myPool.Get() // <1>
	myPool.Put(instance)     // <2>
	myPool.Get()             // <3>
}
