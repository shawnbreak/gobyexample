package main

import (
	"fmt"
	"time"
)

// A comon requirement in programs is  getting the number of seconds,
// milliseconds, nanoseconds since the Unix epoch. Here's how to do it 
// in Go.
func main() {
  // Use time.Now with Unix, UnixMilli or UnixNano to get
  // elapsed time since the Unix epoch in seconds, milliseconds
  // or nanoseconds, respectively.
  now := time.Now()
  fmt.Println(now)

  fmt.Println(now.Unix())
  fmt.Println(now.UnixMilli())
  fmt.Println(now.UnixNano())

  // You can alse convert integer seconds or nanoseconds
  // since epoch into the corresponding time.
  fmt.Println(time.Unix(now.Unix(), 0))
  fmt.Println(time.Unix(0, now.UnixNano()))

}
