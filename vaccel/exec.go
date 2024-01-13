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

	cread := read.cList.list
	cwrite := write.cList.list

	cNrRead := C.ulong(read.cList.size)
	cNrWrite := C.ulong(write.cList.size)

	cRet := C.vaccel_exec_with_resource(&csess, &cobj, cfunc, cread, cNrRead, cwrite, cNrWrite) //nolint:gocritic

	return int(cRet)

}
