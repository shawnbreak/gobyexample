package main

import "fmt"

func main() {

  // range iterates over elements in a variaty of data
  // structures. Let's see how to use range with some
  // of the data structures we've already learned
  nums := []int{2, 3, 4}

  sum := 0
  for _, num := range nums {
    sum += num
  }

  // ranges on arrays and slices provides both the index and value
  // for each entry. Above we didn't need the index, so we ignored it
  // with blank identifier _ . Sometimes we actually want the indexes 
  // though
  fmt.Println(sum)

  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  // ranges on map iterates over key/value pairs
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n:", k, v)
  }

  for k := range kvs {
    fmt.Println("key:", k)
  }


}
