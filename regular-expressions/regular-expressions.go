package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go offers built-in supports for regular expressions.
// Here are some examples of common regexp-related tasks in Go.

func main() {
  // This tests whether a pattern matches a string.
  match, _ := regexp.MatchString("p[a-z]+ch", "peach")
  fmt.Println(match)

  // Above we used a string pattern directly, but for other regexp
  // tasks you'll need to Compile an optimized Regexep struct
  r, _ := regexp.Compile("p[a-z]+ch")

  // Many method are avaialbe on these structs. 
  // Here's a match test like we saw earlier.
  fmt.Println(r.MatchString("peach"))

  // This finds the match for the regexp.
  fmt.Println(r.FindString("peach punch"))

  // This also finds the first match but returns the start and end indexes
  // for the match instead of the matching text.
  fmt.Println("idx:", r.FindStringIndex("peach punch"))

  // The Submatch variants include information about both the whole
  // -pattern matches and submatches within those matches. For examples
  // this will return information for both p([a-z]+)ch and ([a-z]+).
  fmt.Println(r.FindStringSubmatch("peach punch"))

  // Similarly this will return information about the indexes
  // of matches and submatches.
  fmt.Println(r.FindStringSubmatchIndex("peach punch"))

  // The All variants of these functions apply to all matches in the input,
  // not just the first. For example to find all matches in regexp.
  fmt.Println(r.FindAllString("peach punch pinch", -1))

  // These All variants are available for the other functions we
  // saw above.
  fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

  // Providing a non-negative integer as second argument to these functions
  // will limit the number of matches.
  fmt.Println(r.FindAllString("peach punch pinch", 2))

  // Our example above had string arguments and used names like 
  // MatchString. We can alse provide []byte arguments
  // and drop String from the function name. 
  fmt.Println(r.Match([]byte("peach")))

  
  // When creating global variables with regular expressions
  // you can use the MustCompile variation of Compile.
  // MustCompile panics instead of returning an error,
  // Which makes it safer to use for global variables.
  r = regexp.MustCompile("p[a-z]+ch")
  fmt.Println("regexp:", r)

  // The regexp package can alse be used to replace subsets
  // of strings with other values.
  fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

  in := []byte("a peach")

  // The Func variant allows you to transform matched text with 
  // a given function.
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  fmt.Println(string(out))
}
