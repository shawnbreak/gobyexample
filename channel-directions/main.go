package main

import "fmt"


// When using channel as function parameters,
// you can specify if a channelis meant to only send or receive
// values. This specificity increase the type safety of program.

// This ping using only accepts channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
  pings <- msg
}


// This pong function accetps one channel for receives(pings)
// and a second for sends(pongs)
func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}
