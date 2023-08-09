package main

/*

#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L/usr/local/lib -lvaccel -Wl,-rpath=/usr/local/lib -ldl
#include <vaccel.h>
#include <stdlib.h>
#include <stdio.h>

void myPrintFunction2() {
	printf("Hello from inline C\n");
}

typedef struct vaccel_session mysesstype;
typedef unsigned int uinttype;
*/
import "C"

import (
	"fmt"
	//"unsafe"
)

func main() {

	fmt.Println("-------------------------------")

	// C Library
	session := C.mysesstype{};
	flags := 0;
	C.vaccel_sess_init(&session, C.uint32_t(flags));
	C.vaccel_noop(&session);
	C.vaccel_sess_free(&session);
	//C.free(unsafe.Pointer(mystr))

	// Inline C
	//C.myPrintFunction2()

	//fmt.Println("-------------------------------")
}
