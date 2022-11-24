package main

import (
	"fmt"
	"time"
)

func main() {
	// First we’ll look at basic rate limiting. Suppose we want
	// to limit our handling of incoming requests.
	// We’ll serve these requests off a channel of the same name
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("requests", req, time.Now())
	}
}
