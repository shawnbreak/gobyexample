package main

import (
	"fmt"
	"os"
)

// Go offers excellent support for string formatting
// in the printf  tradition. Here are some examples of 
// some common string formatting tasks.

type point struct {
  x, y int
}

func main() {

  p := point{1, 2}

  // Go offer several printing verbs designed to format
  // of our point struct.
  // general Go valus. For example, this prints and instance
  fmt.Printf("struct1: %v\n", p)

  // If the value is struct, the %+v variant will include 
  // the struct's fields.
  fmt.Printf("struct2: %+v\n", p)

  // The %#v variant prints a Go syntax representation of value,
  // i.e. the source code snippt that would produce that value.
  fmt.Println("struct3: %#v\n", p)

  // To print the type of a value, use %T.
  fmt.Printf("type: %T\n", p)

  // Formatting booleans is straight-forward.
  fmt.Printf("bool: %t\n", true)

  // There are many options for formatting integers.
  // Use %d for standard, base-10 formatting.
  fmt.Printf("int: %d\n", 123)

  // This prints a binary representation.
  fmt.Printf("bin: %b\n", 14)

  // This prints character corresponding to the given integer.
  fmt.Printf("char: %c\n", 33)

  // %x provides hex encoding.
  fmt.Printf("hex: %x\n", 456)

  // There are also several formatting options for floats.
  // For basic decimal formatting, use %f.
  fmt.Printf("float1: %f\n", 78.9)

  // %e and %E format the float in (slightly different version of) scientific notion.
  fmt.Printf("float2: %e\n", 1234000000000.0)
  fmt.Printf("float3: %E\n", 1234000000000.0)

  // For basic string printing using %s.
  fmt.Printf("str1: %s\n", "\"string\"")
  
  // To double-quote string s as in Go source, use %q.
  fmt.Printf("str2: %q\n", "\"string\"")

  // As with integers seen earlier, %x renders the string in base-16,
  // with two output characters per byte of input.
  fmt.Printf("str3: %x\n", "hex this")

  // To print a representation of pointer, use %p.
  fmt.Printf("pointer: %p\n", &p)

  // When formatting numbers you will often want to control
  // the width and precision of resulting figure.
  // To specify the width of an integer, use a number after %
  // in the verb. By default, the result will right justified and 
  // padding with spaces.
  fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

  // You will also specify the width of printed floats,
  // though usually you'll also want to restrict the decimal
  // precision at the same time with width.precision syntax.
  fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
  
  // To left-justify, use - flat.
  fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

  // You may also want to control width when formatting strings,
  // especially to ensure that align in table like  output.
  // For basic right-justified width.
  fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

  // To left-justify, use the - flag.
  fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

  // So far we've seen Printf, which prints the formatted string
  // to os.Stdout. Sprintf formats and return a string without printing
  // it anyway.
  s := fmt.Sprintf("sprintf: a %s", "string")

  fmt.Println(s)

  // You can format + print to io.Writers other than os.Stdout
  // using Fprintf.
  fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
