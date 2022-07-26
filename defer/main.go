package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later
// in a program's execution, usually for purpose of cleanup.
// defer is often used where e.g. ensure and finally would be 
// used in other language.

// Suppose we wanted to create a file, write to it, and then close it.
// when we're done. Here's how we could do that with defer.
func main() {

  // Immediately after getting a file object with creaFile,
  // we defer the closing of that file with closeFile.
  // This wil be executed at the end of the enclosing  function(main),
  // after writeFile has finished.
  f := createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)
}

func createFile(s string) *os.File {
  fmt.Println("creating")
  f, err := os.Create(s)
  if err != nil {
    panic(err)
  }
  return f
}

func writeFile(f *os.File) {
  fmt.Println("writing")
  fmt.Fprintln(f, "data")
}

// It's important to check for errors when closing a file,
// even in a defered function.
func closeFile(f *os.File) {
  fmt.Println("closing")
  err := f.Close()
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %v\n", err)
    os.Exit(1)
  }
}



