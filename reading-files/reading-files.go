package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading and writing files are basic tasks for
// for many Go programs. First we'll look at some
// examples of reading files.

// Reading file requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
  if e != nil {
    panic(e)
  }
}


func main() {
  // Perhaps the most basic file reading task
  // is slurping a file's entire contents into memory.
  dat, err := os.ReadFile("/tmp/dat")
  check(err)
  fmt.Println(string(dat))

  // You'll often want to more control over how and
  // and what parts of a file are read. For these task,
  // start by Opening a file to obtain an os.File
  // value.
  f, err := os.Open("/tmp/dat")
  check(err)

  // Read some bytes from the beginning of the file.
  // Allow up to 5 to be read but alse note how many
  // actually were read.
  b1 := make([]byte, 5)
  n1, err := f.Read(b1)
  check(err)
  fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

  // You can also seek to a known location
  // in the file and Read from here.
  o2, err := f.Seek(6, 0)
  check(err)
  b2 := make([]byte, 2)
  n2, err := f.Read(b2)
  check(err)

  fmt.Printf("%d bytes @ %d:", n2, o2)
  fmt.Printf("%v\n", string(b2[:n2]))

  
  // This io package provides some functions that
  // may be helpful for file reading. For example,
  // read like the ones above can be more robustly 
  // implemented with ReadAtLeast.
  o3, err := f.Seek(6, 0)
  check(err)
  b3 := make([]byte, 2)
  n3, err := io.ReadAtLeast(f, b3, 2)
  check(err)

  fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

  // There is no built-in rewind, but Seek(0.0) 
  // accomplishes this.
  _, err = f.Seek(0, 0)
  check(err)

  // This bufio package implements a buffered reader
  // that may be useful both for its efficiency with 
  // many small reads and because of additional reading methods
  // it provides.
  r4 := bufio.NewReader(f)
  b4, err := r4.Peek(5)
  check(err)

  // Close the file when you're done(Usually this would be 
  // be scheduled immediately after Opening with defer),
  fmt.Printf("5 bytes: %s\n", string(b4))
  f.Close()
}
