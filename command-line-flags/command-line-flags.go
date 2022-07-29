// Command-line flags are a common way to specify options
// for command-line programs. For example, in wc -l the -l is 
// a command-line flag.

package main

// Go provides a flag package supporting basic command-line flag parsing.
// We'll use this package to implement our example command-line program.
import (
	"flag"
	"fmt"
)

func main() {
  // Basic flag declaration are available for string, integer, and bool
  // options. Here we declare a string flag word with default value "foo"
  // ans a short description. This flat.String function returns a string pointer
  // (not a sttring value); we'll see how to use this pointer below.
  wordPtr := flag.String("word", "foo", "a string")

  // This declares numb and fork flags, using similar approach to the word flag.
  numbPtr := flag.Int("numb", 42, "an int")
  forkPtr := flag.Bool("fork", false, "a bool")

  // It's alse possible to declare an option that uses existing var
  // declared elsewhere in the program.
  // Note that we need pass in a pointer to the flag declaration function.
  var svar string
  flag.StringVar(&svar, "svar", "bar", "string var")

  // Once all flags are declared, call flag.Parse() to execute the command-line
  // parsing.
  flag.Parse()

  // Here we'll just dump out the parsed options and any trailing positional
  // arguments. Note that we need to dereference the pointers with e.g. *wordPtr
  // to get the actual option values.
  fmt.Println("word:", *wordPtr)
  fmt.Println("numb:", *numbPtr)
  fmt.Println("fork:", *forkPtr)
  fmt.Println("svar:", svar)
  fmt.Println("tail:", flag.Args())
}
