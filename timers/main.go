package main

import (
	"fmt"
	"time"
)

// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// timer and ticker features make both of these takes easy
// We'll look first at timers and then tickers.

func main() {
  // Timer represent a single event in the future.
  // You tell the time how long you want to wait,
  // and it provides a channelq that will be notified at that 
  // time. This timer will will wait 2 seconds.
  timer1 := time.NewTimer(2 * time.Second)

  // The <-timer1.C blocks on the timer's channel C until
  // it sends a value indicating the timer fired.
  <-timer1.C
  fmt.Println("Timer1 fired")

  // If you just want to wait, you could have used time.Sleep(). One
  // reason a timer may be useful is that you can cancel the timer
  // before it fires. Here's is example of that.
  timer2 := time.NewTimer(time.Second)
  go func() {
    <-timer2.C
    fmt.Println("Timer2 fired")
  }()

  stop2 := timer2.Stop()
  if stop2 {
    fmt.Println("Timer 2 stopped")
  }

  // Give thee timer2 enough time to fire, if it ever was going to,
  // to show it is in fact stopped.
  time.Sleep(2 * time.Second)
}
