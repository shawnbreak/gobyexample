package main

import "fmt"

// Go structs are typed collections on fields
// They are useful for groupind data together
// to form records

// This person struct has name and age fields
type person struct {
  name string 
  age int
}

// newPerson constructs a new person struct with
// the given name
func newPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  // You can safely return a pointer to a local variable
  // as a local variable will survive the scope of the function
  return &p
}

func main() {
  // This syntax creates a new struct
  fmt.Println(person{"Bob", 20})

  // You can name the fields when initializing a struct
  fmt.Println(person{name: "Alice", age: 30})

  // Omited fields will be zero valued
  fmt.Println(person{name: "Fred"})

  // An & prefix yeilds a pointer to the struct
  fmt.Println(&person{name: "Ann", age: 40})

  // It's idiomatic to encapsulate new struct creation
  // in construncter function
  fmt.Println(newPerson("Jon"))

  // Access struct fields with a dot
  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)


  // You can also use dots with struct pointers
  // The pointers are automatically dereferenced
  sp := &s
  fmt.Println(sp.age)

  // Structs are mutable
  sp.age = 51
  fmt.Println(sp.age)
  fmt.Println(s.age)
}
