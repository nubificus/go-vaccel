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

*/
import "C"

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------------")
	var session C.struct_vaccel_session
	flags := 0
	e := C.vaccel_sess_init(&session, C.uint32_t(flags))
	if e != 0 {
		fmt.Println("Session not initialized")
	}
	e = C.vaccel_noop(&session)
	if e != 0 {
		fmt.Println("Session not initialized")
	}
	e = C.vaccel_sess_free(&session)
	if e != 0 {
		fmt.Println("Session not initialized")
	}
}
