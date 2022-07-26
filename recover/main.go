package main

import "fmt"

// Go makes is possible to recover from a panic,
// by using the recover built-in function. A recover can stop
// a panic from aborting the program and let it continue with 
// execution instead.

// An example of where this can be useful: a server wouldn't want to
// crash if one of the client connections exhibits a critical error.
// Instead, the server would want to close that connection and continue
// serving other clients. In fact, this is what Go's net/http does
// by default for HTTP servers.


// This function panics.
func myPanic() {
  panic("a problem")
}

func main() {
  // recover must be call with in a deferred function.
  // When the enclosing function panics  the defer will activate 
  // and recover call within it will catch the panic.
  defer func() {
    // The return value of recover is the error raised in the call to panic
    if r := recover(); r != nil {
      fmt.Println("Recovered. Error:\n", r)
    }
  }()
  myPanic()

  // The code will not run, because myPanic panics. The execution of main
  // stops at the point of panic and resumes in deferred closure.
  fmt.Println("after my panic")
}
