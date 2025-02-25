package vaccel

/*
#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

import (
	"unsafe"
)

type Arg struct {
	cArg *C.struct_vaccel_arg
}

type ArgList struct {
	cList *C.struct_vaccel_arg_list
}

/* Type of function to serialize a structure */
/* Returns pointer to serialized data and the size in bytes */
type Serializer func(buf unsafe.Pointer) (unsafe.Pointer, uint32)

/* Type of function to deserialize a structure */
/* Returns pointer to the constructed structure */
type Deserializer func(buf unsafe.Pointer) unsafe.Pointer

func ArgsInit(size uint32) *ArgList {
	list := new(ArgList)
	list.cList = C.vaccel_args_init(C.uint(size))

	return list
}

func (arglist *ArgList) AddSerialArg(buf unsafe.Pointer, size uintptr) int {
	return int(C.vaccel_add_serial_arg(arglist.cList, buf, C.uint(size)))
}

func (arglist *ArgList) AddNonSerialArg(nonSerialBuf unsafe.Pointer,
	argtype uint32, serialize Serializer) int { //nolint:revive // argtype will be used in a next iteration

	serialBuf, bytes := serialize(nonSerialBuf)

	return arglist.AddSerialArg(serialBuf, uintptr(bytes))
}

func (arglist *ArgList) ExpectSerialArg(buf unsafe.Pointer, size uintptr) int {
	return int(C.vaccel_expect_serial_arg(arglist.cList, buf, C.uint(size)))
}

func (arglist *ArgList) ExpectNonSerialArg(expectedSize uintptr) int {
	return int(C.vaccel_expect_nonserial_arg(arglist.cList, C.uint(expectedSize)))
}

func (arglist *ArgList) GetArgs() *Arg {
	args := new(Arg)
	args.cArg = arglist.cList.list

	return args
}

func (args *Arg) ExtractSerialArg(idx int) unsafe.Pointer {
	return C.vaccel_extract_serial_arg(args.cArg, C.int(idx))
}

func (arglist *ArgList) ExtractSerialArg(idx int) unsafe.Pointer {
	return C.vaccel_extract_serial_arg(arglist.cList.list, C.int(idx))
}

func (arglist *ArgList) ExtractNonSerialArg(idx int, deserialize Deserializer) unsafe.Pointer {
	nonSerialBuf := arglist.ExtractSerialArg(idx)
	return deserialize(nonSerialBuf)
}

func (arglist *ArgList) Delete() int {
	return int(C.vaccel_delete_args(arglist.cList))
}
