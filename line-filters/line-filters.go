package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program
// that reads input from stdin, processes it,
// and then prints some derived result to stdout.
// grep and sed are common line filters.

// Here's an example line filter in Go that writes
// a capitlized version of all input text. You can use
// this pattern to write your own Go line filters.
func main() {

  // Wrapping the unbuffered os.Stdin with buffered
  // Scanner gives us convenient Scan method that advances
  // the scanner to next token; which is the next line
  // in the default scanner.
  scanner := bufio.NewScanner(os.Stdin)

  // Text returns the current token, here the next line,
  // from the input.
  for scanner.Scan() {

    ucl := strings.ToUpper(scanner.Text())
    // Wrote out the uppercased line.
    fmt.Println(ucl)
  }

  // Check for errors during Scan. End of file
  // is expected and not reported by Scan as an Error.
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "error", err)
    os.Exit(1)
  }
  
}
