package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities
// for binary or text blobs. For example, TLS/SSL certificates
// use SHA256 to compute a certificate's signature. 
// Here's how to compute SHA256 in Go.


// Go implements several hash functins in various cyrpto/* packages.

func main() {
  s := "sha256 this string"

  // Here we start with a new hash.
  h := sha256.New()
  // Write expects bytes. If you have a string s, use []bytes(s) 
  // to coere it to bytes.
  h.Write([]byte(s))
  // This gets the  finalized hash result asa a byte slice.
  // The argument to Sum can be used to append an existing byte slice:
  // it usually isn't needed.
  bs := h.Sum(nil)

  fmt.Println(s)
  fmt.Printf("%x\n", bs)
}
