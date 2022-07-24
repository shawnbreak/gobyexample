package main

import (
  "fmt"
  "time"
)
// A goroutine is a lightweight thread execution.

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

// Suppose we have function call f(s). Here's how we'd call that in the usual way, runing it synchronously.

func main() {
  f("direct")

   // to invoke this function in a goroutine, use go f(s). this new goroutine will execute concurrently with the calling one
  go f("goroutine")

  // You can alse starts a goroutine for anonymous function call.
  go func(msg string) {
    fmt.Println(msg)
  }("going")

  // Our two functions calls are running asynchronously in seperate goroutines now. Wait for them to finish (for a more robut approach, use a WaitGroup).
  time.Sleep(time.Second)
  fmt.Println("done")


  // When we use this program we see the output of the blocking call first, then the output of the two goroutines. The goroutines' output may be interleaved, because goroutines are being run concurrently by Go routine.
}

