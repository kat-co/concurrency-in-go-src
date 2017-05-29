package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // <1>
		for i := 0; i < b.N; i++ {
			c <- token // <2>
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // <1>
		for i := 0; i < b.N; i++ {
			<-c // <3>
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // <4>
	close(begin)   // <5>
	wg.Wait()
}
