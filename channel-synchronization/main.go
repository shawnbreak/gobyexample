package main

import (
  "fmt"
  "time"
)
 
// We can use channels to synchronize execution across
// goroutines. Here's an example of using blocking receive
// to wait for a goroutine to finish. When waiting for multiple
// goroutines to finish,  you prefer to use WaitGroup.

// This is the function we'll return a goroutine. The done channel
// will be used to notify another goroutine  that this function's worker
// is done.
func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Print("done")

  // Send a value to notify that we are done.
  done <- true
}

func main() {
  // Start a work goroutine, giving it the channel to notify on.
  done := make(chan bool, 1)
  go worker(done)

  // Block until we receive a notification from worker on the channel.
  <- done
}
