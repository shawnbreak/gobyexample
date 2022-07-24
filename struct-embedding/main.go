package main

import "fmt"

// Go supports embedding structs and interfaces
// to express a more seamless composition of types
// This is not to be confused with // go:embed which
// is a go directive introduced in Go version 1.16+
// to embed files and folders into application binary.

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like
// a field without a name.
type container struct {
  base
  str string
}

func main() {
  // When creating structs with literals, we have
  // to initialize the embedding explicitly; 
  // here embedded types serves as the field name.
  co := container{
    base: base{
      num: 1,
    },
    str: "some name",
  }

  // We access the base field directly on co, e.g. co.num.
  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

  // Alternatively, we can spell out the full path using
  // the embedded type name
  fmt.Println("also num:", co.base.num)

  // Since container embeded base, the methods of base
  // also become method of container, here we call a method
  // that was embedded from base directly on co.
  fmt.Println("describe:", co.describe())

  type describer interface {
    describe() string
  }

  // Embedding structs with methods used to bestow interface
  // implementations onto other structs.
  // Here we see that a container now implements the describer 
  // interface because it embed base.
  var d describer = co
  fmt.Println("desciber:", d.describe())
}


