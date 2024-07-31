package main

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	cStr := C.CString("Hello World")
	C.puts(cStr)
	C.free(unsafe.Pointer(cStr))
}
