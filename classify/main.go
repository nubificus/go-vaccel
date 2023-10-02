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
	"unsafe"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./classify <filename>")
		return
	}

	// Get the filename from command line argument
	filePath := os.Args[1]

	imageBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	// Convert the Go byte slice to a C array
	cImageBytes := (*C.uchar)(&imageBytes[0])

	// C Library
	session := C.mysesstype{}
	flags := 0
	cText := (*C.uchar)(C.malloc(C.size_t(256)))
	defer C.free(unsafe.Pointer(cText)) // Free the memory when done
	cOutImageName := (*C.uchar)(C.malloc(C.size_t(256)))
	defer C.free(unsafe.Pointer(cOutImageName)) // Free the memory when done

	e := C.vaccel_sess_init(&session, C.uint32_t(flags))
	if e != 0 {
		fmt.Println("session not created")
		os.Exit(int(e))
	}
	e = C.vaccel_image_classification(&session, unsafe.Pointer(cImageBytes), cText, cOutImageName, C.ulong(len(imageBytes)), C.ulong(256), C.ulong(256))
	if e != 0 {
		fmt.Println("image classification failed")
		os.Exit(int(e))
	}
	e = C.vaccel_sess_free(&session)
	if e != 0 {
		fmt.Println("session not freed ")
		os.Exit(int(e))
	}
}
