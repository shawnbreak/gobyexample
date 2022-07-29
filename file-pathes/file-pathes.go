package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// The filepath package provides functions to parse and construct
// file paths in a way that is portable between operating systems,
// dir/file on Linux vs dir\file on Windows, for example.

func main() {
  // Join should be used construct path in a portable way.
  // It takes any number of arguments and constructs a 
  // hierachical path from them.
  p := filepath.Join("dir1", "dir2", "filename")
  fmt.Println("p:", p)

  // You should alway use Join instead of concatenating /s or \s 
  // manually. In addition to providing protability, Join will also 
  // normalize pathes rmoving superfluous seperators and directory changes
  fmt.Println(filepath.Join("dir1//", "filename"))
  fmt.Println(filepath.Join("dir1/../dir1", "filename"))

  // Dir and Base can be used to split a path to the directory and the file.
  // Alternatively, Split wil return both the same call.
  fmt.Println("Dir(p)", filepath.Dir(p))
  fmt.Println("Base(p)", filepath.Base(p))

  // We can check whether a path is absolute.
  fmt.Println(filepath.IsAbs("dir/file"))
  fmt.Println(filepath.IsAbs("/dir/file"))

  // Some file names has extensions following a dot. We can
  // split the extension out of such names with Ext.
  filename := "config.json"

  ext := filepath.Ext(filename)
  fmt.Println(ext)

  // To find the files name with extension removed, use string.TrimSuffix.
  fmt.Println(strings.TrimSuffix(filename, ext))

  // Rel finds a relative path between a base and a target.
  // It returns an error if the target cannot be made relative to base.
  rel, err := filepath.Rel("a/b", "a/b/t/file")
  if err != nil {
    panic(err)
  }
  fmt.Println(rel)


  rel, err = filepath.Rel("a/b", "a/c/t/file")
  if err != nil {
    panic(err)
  }
  fmt.Println(rel)
}



