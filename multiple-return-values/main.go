package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func main() {

  // to return both result and error values from  function
  // Go has builtin support for multiple return values
  // This feature is used often in idiomatic Go, for example
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
}
