package main

import (
	// char *hello_message();
	"C"
	"fmt"
)

func main() {
	fmt.Println(C.GoString(C.hello_message()))
	// fmt.Println("Hello, world!")
}
