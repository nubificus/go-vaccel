package vaccel

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



func ExecWithResource(sess *Session, obj *SharedObject, funcname string,
	read *ArgList, write *ArgList) int {
	
	csess := sess.c_sess
	cobj  := obj.c_obj
	cfunc := C.CString(funcname)

	cread  := read.c_list.list
	cwrite := write.c_list.list

	c_nr_read  := C.ulong(read.c_list.size)
	c_nr_write := C.ulong(write.c_list.size)

	c_ret := C.vaccel_exec_with_resource(
		&csess, &cobj, cfunc, cread, c_nr_read, cwrite, c_nr_write)

	return int(c_ret)

}
