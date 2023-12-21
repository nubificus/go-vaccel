package vaccel


/*

#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L/usr/local/lib -lvaccel -Wl,-rpath=/usr/local/lib -ldl
#include <vaccel.h>
#include <stdlib.h>
#include <stdio.h>


typedef struct vaccel_session mysesstype;
typedef unsigned int uinttype;
*/
import "C"

import (
	"unsafe"
)

type Arg struct 
{
	c_arg *C.struct_vaccel_arg
}

type ArgList struct 
{
	c_list *C.struct_vaccel_arg_list
}

/* Type of function to serialize a structure */
/* Returns pointer to serialized data and the size in bytes */
type Serializer func(buf unsafe.Pointer) (unsafe.Pointer, uint32)


/* Type of function to deserialize a structure */
/* Returns pointer to the constructed structure */
type Deserializer func(buf unsafe.Pointer) unsafe.Pointer


func ArgsInit(size uint32) *ArgList {

	list := new(ArgList)
	list.c_list = C.vaccel_args_init(C.uint(size))

	return list

}

func (arglist *ArgList) AddSerialArg(buf unsafe.Pointer, size uintptr) int {

	return int(C.vaccel_add_serial_arg(arglist.c_list, buf, C.uint(size)))

}

func (arglist *ArgList) AddNonSerialArg(nonSerialBuf unsafe.Pointer, 
	argtype uint32, serialize Serializer) int {

	serialBuf, bytes := serialize(nonSerialBuf)

	return arglist.AddSerialArg(serialBuf, uintptr(bytes))

}

func (arglist *ArgList) ExpectSerialArg(buf unsafe.Pointer, size uintptr) int {
	
	return int(C.vaccel_expect_serial_arg(arglist.c_list, buf, C.uint(size)))

}

func (arglist *ArgList) ExpectNonSerialArg(expectedSize uintptr) int {
	
	return int(C.vaccel_expect_nonserial_arg(arglist.c_list, C.uint(expectedSize)))

}

func (arglist *ArgList) GetArgs() *Arg {
	
	args := new(Arg)
	args.c_arg = arglist.c_list.list
	
	return args

}

func (args *Arg) ExtractSerialArg(idx int) unsafe.Pointer {
	
	return C.vaccel_extract_serial_arg(args.c_arg, C.int(idx))

}

func (arglist *ArgList) ExtractSerialArg(idx int) unsafe.Pointer {
	
	return C.vaccel_extract_serial_arg(arglist.c_list.list, C.int(idx))
	
}

func (arglist* ArgList) ExtractNonSerialArg(idx int, deserialize Deserializer) unsafe.Pointer {
	
	nonSerialBuf := arglist.ExtractSerialArg(idx)

	return deserialize(nonSerialBuf)

}

func (arglist *ArgList) Delete() int {

	return int(C.vaccel_delete_args(arglist.c_list))

}
