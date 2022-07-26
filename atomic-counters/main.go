package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication
// over channels. We saw this for example with work pools. There are
// a few other options for managing state through. Here we'll look at 
// using the sync/atomic package for atomic counters accessed by multiple
// goroutines

func main() {
  // We'll use an unsigned integer to represent our counter.
  var ops uint64

  // A wait group will help us wait for all goroutines to finish their work.
  var wg sync.WaitGroup

  // We'll start 50 goroutines that each increment the counter exactly 
  // 1000 times.
  for i := 0; i < 50; i++ {
    wg.Add(1)

    // To automatically increment the counter we use AddUint64,
    // giving it the memory address of our ops counter with & syntax.
    go func()  {
      for c := 0; c < 1000; c++ {
        atomic.AddUint64(&ops, 1)
      }
      wg.Done()
    }()
  }

  // Wait until the goroutines are done.
  wg.Wait()

  // It's safe to access ops now because we know no other goroutines
  // is writing to it. Reading atomics safely while they are being updated
  // is also possible, using functions like atomic.LoadUint64().
  fmt.Println("ops:", ops)
}
