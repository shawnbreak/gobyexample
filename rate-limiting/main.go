package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resources
// utilization and maintaining quality of service. Go elegantly supports
// rate limiting with goroutines, channels and ticks.

func main() {
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)

  // First limiter channel will receive a value every 200 milliseconds.
  // This is the regulator in our rate limiting scheme.
  limiter := time.Tick(200 * time.Millisecond)

  // By blocking on a receive from the limiter channel before serving
  // each request, we limit ourselves to 1 request every 200 milliseconds.
  for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
  }


  // We may want to short bursts of requests in our rate limiting scheme
  // while preserving the overall rate limit. We can accomplish this by 
  // buffer our limit channel. This burstyLimiter channel will allow
  // bursts of up 3 events.
  burstyLmiter := make(chan time.Time, 3)

  // Fill up the channel to represent allowed bursting.
  for i := 0; i < 3; i++ {
    burstyLmiter <- time.Now()
  }


  // Every 200 milliseconds we'll try to add a new value to burstyLimiter
  // up to it's limit 3.
  go func() {
    for t := range time.Tick(200 * time.Millisecond) {
      burstyLmiter <- t
    }
  }()


  // Now simulate 5 more incoming requests. The first 3 of these will
  // benefit from burst capability of burstyLimiter.
  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)

  for req := range burstyRequests {
    <-burstyLmiter 
    fmt.Println("request", req, time.Now())
  }

}
