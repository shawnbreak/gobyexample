package main

import "fmt"

func MapKeys[K comparable, V any](m maps[K]V) []K {
  r := make([]K, 0, len(m))
  for k := range m {
    r = append(r, k)
  }
  return r
}

type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val T
}

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
