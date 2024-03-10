package main

/*
#cgo LDFLAGS: ./liblib.a
#include "./lib.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Println(C.GoString(C.hello_message()))
}
