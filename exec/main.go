package main

import (
	"github.com/nubificus/go-vaccel/vaccel"
	"strconv"
	"os"
	"fmt"
	"unsafe"
	"C"
)

func main() {

	/* Read User Args */
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./exec <filename> <input>")
		return
	}

	filePath := os.Args[1]
	input    := os.Args[2]

	inputInt, e := strconv.Atoi(input)
	if e != nil {
		fmt.Println("error converting input")
	}


	/* Shared Object */
	var sharedObject vaccel.SharedObject

	err := vaccel.SharedObjectNew(&sharedObject, filePath)

	if err != 0 {
		fmt.Println("error creating shared object")
		os.Exit(int(err))
	}


	/* Session */
	var session vaccel.Session

	err = vaccel.SessionInit(&session, 0)

	if err != 0 {
		fmt.Println("error initializing session")
		os.Exit(int(err))
	}


	/* Register Shared Object - Session */
	res := sharedObject.GetResource()
	err = vaccel.SessionRegister(&session, res)

	if err != 0 {
		fmt.Println("error registering resource with session")
		os.Exit(int(err))
	}


	/* Create the arg-lists */
	read  := vaccel.ArgsInit(1)
	write := vaccel.ArgsInit(1)

	if read == nil || write == nil {
		fmt.Println("Error Creating the arg-lists")
		os.Exit(0)
	}


	/* Add a serialized arg */
	buf  := unsafe.Pointer(&inputInt)
	size := unsafe.Sizeof(inputInt)

	if read.AddSerialArg(buf, size) != 0 {
		fmt.Println("Error Adding Serialized arg")
		os.Exit(0)
	}


	/* Define an expected argument */
	var output int
	buf  = unsafe.Pointer(&output)
	size = unsafe.Sizeof(output)

	if write.ExpectSerialArg(buf, size) != 0 {
		fmt.Println("Error defining expected arg")
		os.Exit(0)
	}


	/* Run the operation */
	err = vaccel.ExecWithResource(&session, &sharedObject, "mytestfunc", read, write)

	if err != 0 {
		fmt.Println("An error occured while running the operation")
		os.Exit(err)
	}


	/* Read the output */
	fmt.Println("Output(1): ", C.uint(output))

	/* Or */
	outbuf := write.ExtractSerialArg(0)

	cast := (*int)(outbuf)
	val  := C.uint(*cast)
	fmt.Println("Output(2): ", val)

	/* Or */
	outbuf = write.GetArgs().ExtractSerialArg(0)

	cast = (*int)(outbuf)
	val  = C.uint(*cast)
	fmt.Println("Output(3): ", val)


	/* Delete the lists */
	if write.Delete() != 0 || read.Delete() != 0 {
		fmt.Println("An error occured in deletion of the arg-lists")
		os.Exit(0)
	}


	/* Free Session */
	if vaccel.SessionFree(&session) != 0 {
		fmt.Println("An error occured while freeing the session")
	}
}
