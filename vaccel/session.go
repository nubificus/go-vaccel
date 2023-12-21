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

type Session struct {
	c_sess C.struct_vaccel_session
}

func SessionInit(sess *Session, flags uint32) int {
	
	return int(C.vaccel_sess_init(&sess.c_sess, C.uint32_t(flags)))

}

func SessionRegister(sess *Session, res *Resource) int {

	return int(C.vaccel_sess_register(&sess.c_sess, res.c_res))

}

func SessionFree(sess *Session) int {

	return int(C.vaccel_sess_free(&sess.c_sess))

}