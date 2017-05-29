package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // <1>
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // <2>
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
