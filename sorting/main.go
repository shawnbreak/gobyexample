package main

import (
	"fmt"
	"sort"
)

// Go's sort package implements sorting for builtins
// user defined types. We'll look at sorting for builtins first.

func main() {

  // Sorts method are specific to buitlin types; here's an example for 
  // strings. Note that sorting is in-place, so it changes the given slice
  // and doesn't return a new one.
  strs := []string{"c", "a", "b"}
  sort.Strings(strs)
  fmt.Println("Strings:", strs)

  // An example of sort ints.
  ints := []int{7, 2, 4}
  sort.Ints(ints)
  fmt.Println("Ints:", ints)

  // We can alse use sort to check if a slice is already is sorted order.
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted:", s)
}
