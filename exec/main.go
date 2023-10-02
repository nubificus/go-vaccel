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
	"os"
	"strconv"
	"unsafe"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: ./exec <filename> <input>")
		return
	}

	// Get the filename from command line argument
	filePath := os.Args[1]

	input := os.Args[2]

	var sharedObject C.struct_vaccel_shared_object

	// int ret = vaccel_shared_object_new(&object, argv[1]);
	fmt.Println("File: ", filePath)
	filePathC := C.CString(filePath)
	err := C.vaccel_shared_object_new(&sharedObject, filePathC)
	if err != 0 {
		fmt.Println("error creating shared object")
		os.Exit(int(err))
	}

	// C Library
	var session C.struct_vaccel_session
	flags := 0
	// ret = vaccel_sess_init(&sess, flags);
	err = C.vaccel_sess_init(&session, C.uint32_t(flags))
	if err != 0 {
		fmt.Println("error initializing session")
		os.Exit(int(err))
	}

	// ret = vaccel_sess_register(&sess, object.resource);
	err = C.vaccel_sess_register(&session, sharedObject.resource)
	if err != 0 {
		fmt.Println("error registering resource with session")
		os.Exit(int(err))
	}

	// Create C pointers for input and output
	// inputC := C.uint(strconv.Atoi(input))
	inputInt, e := strconv.Atoi(input)
	if e != nil {
		fmt.Println("error converting input")
	}
	inputC := C.uint(inputInt)

	var output [200]byte
	// Create a C struct_vaccel_arg array for input and output
	var readArg *C.struct_vaccel_arg
	readArg = (*C.struct_vaccel_arg)(C.malloc(C.ulong(unsafe.Sizeof(*readArg) * 2)))
	argArr := unsafe.Slice((*C.struct_vaccel_arg)(readArg), 2)
	defer C.free(unsafe.Pointer(readArg))
	argArr[0].size = C.uint(len(input))
	argArr[0].buf = (unsafe.Pointer(&inputC))
	argArr[1].size = C.uint(len(output))
	argArr[1].buf = (unsafe.Pointer(&output[0]))

	funcname := C.CString("mytestfunc")
	// ret = vaccel_exec_with_resource(&sess, &object, "mytestfunc", read, 1, write, 1);
	err = C.vaccel_exec_with_resource(&session, (&sharedObject), funcname, &argArr[0], C.ulong(1), &argArr[1], C.ulong(1))
	if err != 0 {
		fmt.Println("error running exec with resource")
		os.Exit(int(err))
	}

	fmt.Printf("output: %s", output)

	err = C.vaccel_sess_free(&session)
	if err != 0 {
		fmt.Println("error cleaning up session")
		os.Exit(int(err))
	}

}
