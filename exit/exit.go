// Use os.Exit to immediately exit with a given status.
package main

import (
	"fmt"
	"os"
)

func main() {
  // defers will not be run when using os.Exit,
  // so this fmt.Println will never be called.
  defer fmt.Println("!")

  // Exit with statu 3.
  os.Exit(3)
}
