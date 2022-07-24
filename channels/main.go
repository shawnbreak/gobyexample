package main

import "fmt"

// Channels are the pipes that connect the concurrent goroutines. You can send values into channels from one goroutine and receive those values from another goroutine.
func main() {

  // Create a new channel with make(chan val-type).
  // Channels are typed by the values they convey.
  messages := make(chan string)

  // Send a value into a channel using the channel <- syntax.
  // Here we send "ping"q to the message channel we made above,
  // from new a goroutine.
  go func() {
    messages <- "ping"
  }()

  // The <- channel syntac receive a value from the channel.
  // Here we receive the "ping" message we send above and print it out.
  msg := <- messages
  fmt.Println(msg)

  // When we run the go program the "ping" message is successfully passed\  // from one goroutine to another via channel.


  // By default sends and receives block until both the receiver and sender are ready.
  // This property allowed us to wait at the end of our program for "ping"
  // message without having to use any other synchronization.
}
