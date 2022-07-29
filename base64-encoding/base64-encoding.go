package main

import (
  // This syntax imports encoding/bas64 package with the b64 name
  // instead of the default base64. It'll save us some space below.
  b64 "encoding/base64"
  "fmt"
)

// Go provides built-in support for base64 encoding/decoding.


func main() {

  // Here's the string we'll encode/decode.
  data := "abc123!?$*&()'-=@~"

  // Go supports both standard and URL-compatible base64.
  // Here's how to encode using standard encoder. The encoder
  // requires a []byte so we convert our string to that type.
  sEnc := b64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(sEnc)

  // Decoding may return an error, which you can check
  // if you don't already know input to be well-formed.
  sDec, _ := b64.StdEncoding.DecodeString(sEnc)
  fmt.Println(string(sDec))
  fmt.Println()

  // This encodes/decodes using a URL-compatible base64 format.
  uEnc := b64.URLEncoding.EncodeToString([]byte(data))
  fmt.Println(uEnc)
  uDec, _ := b64.URLEncoding.DecodeString(uEnc)
  fmt.Println(string(uDec))

}

