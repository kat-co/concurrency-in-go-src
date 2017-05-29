package main

import (
	"fmt"
)

func main() {
	var data int
	go func() { // <1>
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
