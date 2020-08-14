package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callClib/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a c func")
	C.cHello()

	fmt.Println("Going to call another C func")
	myMessage := C.CString("This is Padeba")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All done!")
}