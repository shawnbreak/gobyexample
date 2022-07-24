package main

import "fmt"
// Starting with version 1.18, Go has add support for generics
// also known as type parameters.

// As an example of generic function, MapKeys takes a map of any
// type and returns a slice of it's key. This function has two
// type parameters-K and V; K has comparable constraint, meaning that
// we can compara value with == and != operators. This is required for 
// map keys in Go, V has any contraint, meaning that it's not restrict
// in any value(any is an alias for interface{}).
func MapKeys[K comparable, V any](m maps[K]V) []K {
  r := make([]K, 0, len(m))
  for k := range m {
    r = append(r, k)
  }
  return r
}

// As an example of generic type,
// List is a single-linked list 
// with values of any type
type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val T
}

// We can define methods on generic types 
// just like we do on regular types, but we
// have to keep the type parameters in place.
// The type is List[T], not List
func (list *List[T]) Push(v T) {
  if list.tail == nil {
    list.head = &element[T]{val: v}
    list.tail = list.head
  } else {
    list.tail.next = &element[T]{val: v}
    list.tail = list.tail.next
  }
}

func (list *List[T]) GetAll() []T {
  var elems []T
  for e := list.head; e != list.tail; e = e.next {
    elems = append(elems, e)
  }
  return elems
}


func main() {
  var m = map[int]string{1: "1", 2: "2", 3: "3"}
  fmt.Println("keys m:", MapKeys(m))


  _ = MapKeys[int, string](m)

  list := List[int]{}
  list.Push(10)
  list.Push(13)
  list.Push(20)
  fmt.Println("list:", list.GetAll()) 
}
