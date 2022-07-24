package main

import (
  "fmt"
  "errors"
)

// In Go it's idiomatic to communicate errors 
// via an explicit, seperate return value. This 
// contrasts with exceptions used in other languages
// like java and ruby and the overloaded single result
// / error value sometimes used in C. Go's approach makes
// it easy to see which function return errors and to handle
// them using the same language constructs employed for any others,
// non-error tasks


// By convention, errors are the last return value and hava type error,
// a builtin interface
func f1(arg int) (int, error) {
  if arg == 42 {
    // errors, New constructs a basic error value with the given error message.
    return -1, errors.New("can't work with 42")
  }
  // A nil value in error position indicates that there was no error
  return arg + 3, nil
}


// It's possible to use custom types as error by implementing the Error method on them. Here's a variant on example above that uses a custome type to explicitly  represent an argument error.
type argError struct {
  arg int
  prob string
}

func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// In this case we use &argError syntax to build a new struct,
// supplying values for two fields arg and prob
func f2(arg int) (int, error) {
  if arg == 42 {
    return -1, &argError{arg, "can't work with it"}
  }
  return arg + 3, nil
}

// The two loops below test out each of our error-returning
// functions. Note that we use of an inline error check on if line
// is a common idiom in Go code.
func main() {
  for _, i := range []int{7, 42} {
    if r, e := f1(i); e != nil {
      fmt.Println("f1 failed: ", e)
    } else {
      fmt.Println("f1 worked: ", r)
    }
  }

  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed: ", e)
    } else {
      fmt.Println("f2 worked: ", r)
    }
  }

  // If you want to programmatically use data in a custome error, you'll need to get the error as an instance of the custome error type via type assertion
  _, e := f2(42)  

  if ae, ok := e.(*argError); ok {
      fmt.Println(ae.arg)
      fmt.Println(ae.prob)
  }
}

