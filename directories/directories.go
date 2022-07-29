package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Go has several useful functions for working with directories
// in file system.

func check(e error) {
  if e != nil {
    panic(e)
  }
}


func main() {
  // Create a new sub-directory in current working directory.
  err := os.Mkdir("subdir", 0755)
  check(err)

  // When creating temporay directories, it's good practice to 
  // defer their removeal. os.RemoveAll will delete a whole directory 
  // tree (similarly to rm -rf).
  defer os.RemoveAll("subdir")

  createEmptyFile := func(name string) {
    d := []byte("")
    check(os.WriteFile(name, d, 0644))
  }

  createEmptyFile("subdir/file1")

  // We can create a hierachy diretories, including with parents with 
  // MkdirAll. This is similar to the command line mkdir -p.
  err = os.MkdirAll("subdir/parent/child", 0755)
  check(err)


  createEmptyFile("subdir/parent/file2")
  createEmptyFile("subdir/parent/file3")
  createEmptyFile("subdir/parent/child/file4")

  // ReadDir lists directory contents, returning a slice of
  // os.DirEntry objects.
  c, err := os.ReadDir("subdir/parent")
  check(err)

  fmt.Println("Listijng subdir/parent")
  for _, entry := range c {
    fmt.Println(" ", entry.Name(), entry.IsDir())
  }

  // ChDir lets us change the current working directory, similar to cd.
  err = os.Chdir("subdir/parent/child")
  check(err)

  // Now we'll see the contents of subdir/parent/child when listing 
  // the current directory.
  c, err = os.ReadDir(".")
  check(err)

  fmt.Println("Listijng subdir/parent/child")
  for _, entry := range c {
    fmt.Println(" ", entry.Name(), entry.IsDir())
  }

  // cd back to where we started.
  err = os.Chdir("../../..")
  check(err)

  // We can also visit a directory recursively, including all its
  // sub-directories. Walk accepts a callback function to handle 
  // every file of directory visited.
  fmt.Println("Visiting subdir")
  err = filepath.Walk("subdir", visit)

}

// visit is called for every file or directory founded recursively by filepath.Walk.
func visit(p string, info os.FileInfo, err error) error {
  if err != nil {
    return nil
  }
  fmt.Println(" ", p, info.IsDir())
  return nil
}



