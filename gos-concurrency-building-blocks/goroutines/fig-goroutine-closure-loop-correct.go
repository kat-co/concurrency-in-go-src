package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) { // <1>
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // <2>
	}
	wg.Wait()
}
