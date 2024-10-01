package main

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: -L${SRCDIR}\out -l go_and_rust.dll
#include "./out/go_and_rust.h"
#include <stdio.h>
#include <windows.h>
*/

/*
#cgo LDFLAGS: ./out/libgo_and_rust.a -ldl
#include "./out/go_and_rust.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func main() {
	argsWithoutProg := os.Args[1:]
	word := strings.Join(argsWithoutProg, " ")

	str1 := C.CString(word)
	defer C.free(unsafe.Pointer(str1))

	out := C.count_numbers(str1)

	fmt.Printf("Counted %d numbers\n", out)
}
