package main

import (
	"fmt"
	//"os"
	"unicode/utf8"
)

// A Go string is a readonly slice of bytes.
// The language andq the standard library threat
// string speciallly - as containers of text encoded
// in UTF-8. In other languages, strings are made of
// "characters". In Go, the concept of a character
// is called a rune - it's an integer that represents
// a Unicode code point. This Go blog post is a good
// introduction to the topic

func main()  {
  // s is a string assigned a literal
  // representing the world hello in 
  // the Thai language. Go string literals 
  // are UTF-8 encoded text.
  const s = "สวัสดี"

  // Since string equivalent to []byte, this will produce
  // the length of the raw bytes stored within.
  fmt.Println(len(s))


  // Indexing into a string produces the raw bytes
  // values at each index. This loop generates hex
  // values of all the bytes that constitute the code
  // point in s
  for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
  }

  fmt.Println()

  // To count how many runes are in a string.
  // We can use the utf8 package. Note that the run-time
  // of RuneCountInString depends on the size of the string
  // because it has decode each utf8 rune sequentially.
  // Some Thai characters are represented by multiple UTF-8
  // code points, so the results of this count may be surprising
  fmt.Println("Rune count:", utf8.RuneCountInString(s))

  // A range loop handls strintg specially and decode each rune
  // along with its offset in the string
  for idx, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, idx)
  }
  // os.Exit(1)

  // we can acheive the same iteration by using
  // the utf8.DecodeRuneInString function explicitly
  fmt.Println("\nUsing DecodeRuneInString")
  for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w = width

    // This demonstrates passing a rune value to a funciton
    examineRune(runeValue)
  }
}

// values enclosed in single quotes are rune literal
// we can compare a rune value to a rune literal directly
func examineRune(r rune) {
  if r == 't' {
    fmt.Println("found tee")
  } else if r == 'ส' {
    fmt.Println("found sua")
  }
}
