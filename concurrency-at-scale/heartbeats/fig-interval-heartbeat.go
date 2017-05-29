package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{}) // <1>
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)       // <2>
			workGen := time.Tick(2 * pulseInterval) // <3>

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // <4>
				}
			}
			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse: // <5>
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse: // <5>
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) }) // <1>

	const timeout = 2 * time.Second               // <2>
	heartbeat, results := doWork(done, timeout/2) // <3>
	for {
		select {
		case _, ok := <-heartbeat: // <4>
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results: // <5>
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout): // <6>
			return
		}
	}
}
