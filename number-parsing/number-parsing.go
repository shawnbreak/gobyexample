package main

import (
	"fmt"
	"strconv"
)

// Parsing nubmers from strings is a basic but common
// task in many programs, here's how to it in Go. 

// The default package strconv provides the number parsing.

func main() {

  // With ParseFloat, this 64 tells how many bits of precision to parse.
  f, _ := strconv.ParseFloat("1.234", 64)
  fmt.Println(f)

  // For ParseInt, the 0 means infer the base from the string,
  // 64 requires that the results fits in 64 bits.
  i, _ := strconv.ParseInt("123", 0, 64)
  fmt.Println(i)

  // ParseInt will recognize hex-formatted numbers.
  d, _ := strconv.ParseInt("0x1c8", 0, 64)
  fmt.Println(d)

  // A ParseUint is also available.
  u, _ := strconv.ParseUint("789", 0, 64)
  fmt.Println(u)

  // Atoi is a convenience function for base-10 int parsing.
  k, _ := strconv.Atoi("135")
  fmt.Println(k)

  // Parse functions returns an error on bad input.
  _, e := strconv.Atoi("wat")
  fmt.Println(e)
}
