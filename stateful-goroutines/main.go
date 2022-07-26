package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In the previos example we used explicte locking with mutexes
// to sychronize access to shared state across multiple goroutines
// Another option is to use the built-in sychronization features
// of goroutines and channels to achieve the same result. This
// channel based approach aligns with Go's ideas of sharing memory
// by communicating and having each pecie of data owned by 
// exactly 1 goroutine.


// In this example our state will be owned by a single goroutine.
// This will guarantee that the data is never corrupted with conccurency
// accessd. In order to read or write that state, other goroutines will
// send messages to the owning goroutine and recieve corresponding replies
// These readOp and writeOp structs encapsulate those requests and a way
// for owning goroutine to respond.
type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  val int
  resp chan bool
}

func main() {

  // As before we'll count how many operations we perform
  var readOps uint64
  var writeOps uint64

  // The reads and writes channel will be used by other goroutines
  // to issue read and write requests, respectively.
  reads := make(chan readOp)
  writes := make(chan writeOp)

  // Here is the goroutine that owns the state, which is a map
  // as in  the previos example but now private to the stateful
  // goroutine. This goroutine repeatedly selects on the reads and
  // writes channel, responding to requests as the arrive.
  // A response is executed by first performing the requested operation
  // and then sending a value on the response channel resp to indicate
  // success(and desired value in the case of reads)
  go func() {
    var state = make(map[int]int)
    for {
      select {
      case read := <-reads:
        read.resp <- state[read.key]
      case write := <-writes:
        state[write.key] = write.val
        write.resp <- true
      }
    }
  }()

  // This starts 100 goroutines to issue reads to the state
  // owning goroutine via reads channel. Each read requires constructing
  // a readOp, sending it over the reads channel, and then receiving the
  // result over provided resp channel
  for r := 0; r < 100; r++ {
    go func() {
      for {
        read := readOp{
          key: rand.Intn(5),
          resp: make(chan int),
        }
        reads <- read
        <-read.resp
        atomic.AddUint64(&readOps, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  // We start 10 goroutines as well, using a similar approach.
  for w := 0; w < 10; w++ {
    go func() {
      for {
        write := writeOp{
          key: rand.Intn(5),
          val: rand.Intn(100),
          resp: make(chan bool),
        }

        writes <- write
        <-write.resp
        atomic.AddUint64(&writeOps, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  // Let the goroutines work for a second.
  time.Sleep(time.Second)

  // Finally capture the report op counts.
  readOpsFinal := atomic.LoadUint64(&readOps)
  fmt.Println("readOps:", readOpsFinal)
  writeOpsFinal := atomic.LoadUint64(&writeOps)
  fmt.Println("wirteOps:", writeOpsFinal)
}
