package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use wait group.

// This is the function we'll run in every goroutine.
func worker(id int) {
  fmt.Printf("worker %d starting\n", id)

  // Sleep to simulate a expensive task.
  time.Sleep(time.Second)
  fmt.Printf("worker %d done\n", id)
}

func main() {

  // This waitGroup is used to wait for all the goroutines
  // launched here to finish. Note: if a WaitGroup is explicitly
  // passed into functions, it should be done by pointer.
  var wg sync.WaitGroup


  for i := 1; i <= 5; i++ {
    // Launch several goroutines and increment WaitGroup counter for each.
    wg.Add(1)

    // Avoid reuse of the same i value in each goroutine closure.
    i := i

    // Wrap the worker call in a closure that makes sure to tell 
    // WaitGroup  that this worker is done. This way the worker it self
    // does not have to be aware of the concurrency primitives involved
    // in it's execution.
    go func() {
      defer wg.Done()
      worker(i)
    }()
  }

  // Block until the WaitGroup counter goes back to 0; all the
  // workers nofied they're done.
  wg.Wait()

  // Note that this approach has no straighforward way to propagate errors
  // from workers. For more advanced use case, consider using errgroup pacakge.
}
