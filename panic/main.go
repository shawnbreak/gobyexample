package main

import "os"

// A panic typically means something went unexpectedly wrong.
// Mostly we use it for fail fast on errors that shouldn't occur
// during normal operation, or that we aren't prepared to handle
// gracefully.
func main() {

  // We'll use panic throughout this site to check for unexpected
  // errors. This is ths only program on the site designed to panic.
  panic("a problem")

  // A common use of panic is to abort if a function returns an error
  // value that we don't know how to handle. Here's and example of panic
  // if we get an unexpected error when creating a new file.
  _, err := os.Create("/tmp/file")
  if err != nil {
    panic(err)
  }
}
