// Writing a basic HTTP server is easy using net/http pacakge.

package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in net/http servers is handlers.
// A handler is an object implementing the http.Handler interface.
// A common way to write a handler is by using the http.HandleFunc
// adapter on functions with the appropriate  signature.
func hello(w http.ResponseWriter, req *http.Request) {
  // Functions serve as handlers take a http.ResponseWrite and http.Request
  // as arguments. The response writer is used to fill the HTTP response.
  // Here our simple response just "hello\n"
  fmt.Fprintf(w, "hello\n")
}

// This handler does something a litter more  sophisticated
// by reading all the HTTP request headers and echoing them
// into the response body.
func headers(w http.ResponseWriter, req *http.Request) {
  for name, headers := range req.Header {
    for _, h := range headers {
      fmt.Fprintf(w, "%v: %v\n", name, h)
    }
  }
}

func main() {
  // We register our handler on server routes using http.HandleFunc
  // convenience function. It sets up the default router in the net/http
  // package and takes a function as an argument.
  http.HandleFunc("/hello", hello)
  http.HandleFunc("/headers", headers)


  // Finally we can call the ListenAndServe with the port and the handler.
  // nil tells it to use the default router we've just set up.
  http.ListenAndServe(":8090", nil)
}
