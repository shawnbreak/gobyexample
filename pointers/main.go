package main

import "fmt"

// Go supports pointers
// Allow you to passing references
// to values and records within your program


// We'll show how pointers work in contrast to values 
// with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will 
// be passed to it by value.
// zeroval will get a copy of ival distinct from the one
// in the calling function.

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)

  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}
