// Some command-line tools, like go tool or git have many sub commands
// each with it's own set of flags. For example, go build and go get
// are two different subcommds of go tool. The flag package lets us
// easily define simple subcommands that have their own flags.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
  // We declare a subcommand using NewFlagSet function,
  // and proceed to define new flags specific for this subcommand.
  fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
  fooEnable := fooCmd.Bool("enable", false, "enable")
  fooName := fooCmd.String("name", "", "name")

  // For different subcommand we can define different supported flags.
  barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
  barLevel := barCmd.Int("level", 0, "level")

  // The sub command is expected as the first command to the program.
  if len(os.Args) < 2 {
    fmt.Println("expected 'foo' or 'bar' subcommands")
  }

  // Check wich subcommand is invoked.
  switch os.Args[1] {
  // For every subcommand, we parse it's own flags and have access
  // to trailing positional arguments.
  case "foo":
    fooCmd.Parse(os.Args[2:])
    fmt.Println("subcommand 'foo'")
    fmt.Println("    enable:", *fooEnable)
    fmt.Println("    name:", *&fooName)
    fmt.Println("    tail:", fooCmd.Args())
  case "bar":
    barCmd.Parse(os.Args[2:])
    fmt.Println("subcommand 'bar'")
    fmt.Println("    level:",*barLevel)
    fmt.Println("    tail:",barCmd.Args())
  default:
    fmt.Println("expected 'foo' or 'bar' subcommands")
    os.Exit(1)
  }
}