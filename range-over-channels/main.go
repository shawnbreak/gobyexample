package main

import "fmt"

// In a previous example we saw how for and range provide
// iteration over basic data structures. We can alse use this
// syntax to iterate over values received from a channel.
func main() {


  // We'll iterate over 2 values in the queue channel.
  queue := make(chan string, 2)

  queue <- "one"
  queue <- "two"

  close(queue)

  // This range iterations over each element as its received from 
  // queue. Because we closed the channel abover, the iteration
  // terminates after receiving the 2 elements.
  for elem := range queue {
    fmt.Println(elem)
  }
}


