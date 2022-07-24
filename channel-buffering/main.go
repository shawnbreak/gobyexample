package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept
// sends(chan <-) if there is a corresponding receive(<-chan) 
// ready to receiv the sent value. Buffered channels accept a limited number of values
// without corresponding receiver for those values.
func main() {


  // Here we make a channel of strings buffered up to 2 values.
  messages := make(chan string, 2)

  // Because the channel is buffered, we can send these values into the 
  // channel without a corresponding conccurent receive.
  messages <- "buffered"
  messages <- "channel"

  // Later we can receive those two values as usual.
  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
