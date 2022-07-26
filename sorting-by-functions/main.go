package main

import (
	"fmt"
	"sort"
)

// Sometimes we'll want to sort a collection by something other than
// its natural order. For example, suppose we wanted to sort strings
// by their length instead of alphabetically. Here's an example
// of custom sorts in Go.

// In order to sort by a custome function in Go,
// We need a responding type. Here we've created a byLength type
// that is just an alias for builtin []string type.
type byLength []string

// We implement sort.Interface - Len, Less, ans Swap - on
// our type so we can use the sort package's generic Sort function.
// Len and Swap will usually be similar across types and Less
// will hold the actual costom sorting logic. In our case
// we want to sort in order of increasing string length, 
// so we use len(s[i]) and len(s[j]) here.
func (s byLength) Len() int {
  return len(s)
}

func (s byLength) Swap(i, j int) {
  s[j], s[i] = s[i], s[j]
}

func (s byLength) Less(i, j int) bool {
  return len(s[i]) < len(s[j])
}

// With all of these in place, we can now implement our custom type sort
// by converting original fruits slice ot byLength, and then use sort.Sortsort
// by converting original fruits slice ot byLength, and then use sort.Sort
// on that typed slice.
func main() {
  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(byLength(fruits))
  fmt.Println(fruits)
}
