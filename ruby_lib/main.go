// Compile with :
// go build -buildmode=c-shared -o my_lib.so main.go
// Be aware : Any internal C usage need manual memory free

package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C" // Keep comment and import attached and IN ONE LINE IMPORT

import (
  "unsafe"
  "fmt"
)

//export sayHello
func sayHello() (result *C.char) {
  result = C.CString(SayHello()) // Go String -> C *Char (malloc based on len of content)
  fmt.Println("Memory Allocation for pointer address", result)
  return
}

//export clearChar
func clearChar(pchar *C.char) {
  fmt.Println("Memory Dealocate for pointer address", pchar)
  C.free(unsafe.Pointer(pchar)) // C stdlib:free -> Dealocate memory
}


// Here is internal Go. We return in the comfort of Garbage Collected Runtime
func SayHello() string {
  return "Hello you !"
}

func main() {}
