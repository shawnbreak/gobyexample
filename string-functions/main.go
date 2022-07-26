package main

import (
  "fmt"
  s"strings"
)


// The standart library's string package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

// We alias fmt.Println to a shorter name as we'll use it a lot below.
var p = fmt.Println

func main() {
  // Here's a sample of the functions avaiable in strings.
  // Since these are functions from the pacakge, not methods
  // on string object itself, we need pass the string in question
  // as the first argument to the function. You can find more
  // functions in string  packages docs.
  p("Contains:  ", s.Contains("test","es"))
  p("Count:     ", s.Count("test", "t"))
  p("HasPrefix: ", s.HasPrefix("test", "te"))
  p("HasSuffix: ", s.HasSuffix("test", "st"))
  p("Index:     ", s.Index("test", "e"))
  p("Join:      ", s.Join([]string{"a", "b"}, "-"))
  p("Repeat:    ", s.Repeat("a", 5))
  p("Replace:   ", s.Replace("foo", "o", "0", -1))
  p("Replace:   ", s.Replace("foo", "o", "0", 1))
  p("Splite:    ", s.Split("a-b-c-d-e", "-"))
  p("ToLower:   ", s.ToLower("TEST"))
  p("ToUpper:   ", s.ToUpper("test"))

}
