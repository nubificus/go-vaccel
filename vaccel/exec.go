package vaccel

/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

func ExecWithResource(sess *Session, obj *SharedObject, funcname string,
	read *ArgList, write *ArgList) int {

	csess := sess.cSess
	cobj := obj.cObj
	cfunc := C.CString(funcname)

	cread := read.c_list.list
	cwrite := write.c_list.list

	c_nr_read := C.ulong(read.c_list.size)
	c_nr_write := C.ulong(write.c_list.size)

	c_ret := C.vaccel_exec_with_resource(
		&csess, &cobj, cfunc, cread, c_nr_read, cwrite, c_nr_write)

	return int(c_ret)

}
