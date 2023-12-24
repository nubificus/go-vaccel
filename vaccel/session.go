package vaccel


/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

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
