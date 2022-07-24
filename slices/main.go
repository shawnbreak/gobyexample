package main

import "fmt"

func main() {

  // Slices are a key data type in Go,
  // giving a more powerful interface to sequences than arrays

  // Unlike arrays, slices are typed only by the elements the contain
  // not the number of elements.
  // To create an empty slice with non-zero length, use the builtin make.
  // Here we are make a slice of string of length 3 
  // (initially zero-valued)
  s := make([]string, 3)

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"

  fmt.Println("set:", s)
  fmt.Println("get:", s[2])
  
  // len returns the length of the slice as expected
  fmt.Println(len(s))

  // In addition to these basic operations, slices support 
  // serveral more that make them richer than arrays.
  // One is builtin append, which returns a slice containing 
  // one or more new values. Note that we need to accept 
  // a return value from  append as we may get a new slice value
  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)


  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy", c)

  l := s[2:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  twoD := make([][]int, 3)

  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)
}


