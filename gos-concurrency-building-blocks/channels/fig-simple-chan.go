package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!" // <1>
	}()
	fmt.Println(<-stringStream) // <2>
}
