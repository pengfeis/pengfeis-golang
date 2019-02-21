package main

import (
	"time"
	"fmt"
)

func main()  {

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
	}
	
	go func() {
        for t := range time.Tick(500 * time.Millisecond) {
            burstyLimiter <- t
        }
	}()
	
	burstyRequests := make(chan int, 15)
    for i := 1; i <= 15; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
	
}