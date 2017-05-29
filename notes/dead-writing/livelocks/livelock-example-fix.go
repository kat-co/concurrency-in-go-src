package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		sync.Mutex
		id     string
		locked bool
		value  int
	}

	lock := func(v *value) {
		v.Lock()
		v.locked = true
	}
	unlock := func(v *value) {
		v.Unlock()
		v.locked = false
	}
	printSum := func(wg *sync.WaitGroup, id string, v1, v2 *value) {
		defer wg.Done()
		var sum int
		for i := 0; ; i++ { // <4>
			if i >= 5 {
				fmt.Println("canceling goroutine...")
				return
			}

			fmt.Printf("%v: acquiring lock on %v\n", id, v1.id)
			lock(v1) // <1>

			time.Sleep(2 * time.Second)

			if v2.locked { // <2>
				fmt.Printf("%v: releasing lock on %v\n", id, v1.id)
				unlock(v1) // <3>
				fmt.Printf("%v: %v locked, retrying\n", id, v2.id)
				continue
			}

			fmt.Printf("%v: acquiring lock on %v\n", id, v2.id)
			lock(v2)

			sum = v1.value + v2.value
			fmt.Printf("%v: releasing lock on %v\n", id, v1.id)
			unlock(v1)

			fmt.Printf("%v: releasing lock on %v\n", id, v2.id)
			unlock(v2)
			break
		}

		fmt.Printf("sum: %v\n", sum)
	}
	a, b := value{id: "a"}, value{id: "b"}
	var wg sync.WaitGroup
	wg.Add(2)
	go printSum(&wg, "first", &a, &b)
	go printSum(&wg, "second", &a, &b) // <1>

	wg.Wait()
}
