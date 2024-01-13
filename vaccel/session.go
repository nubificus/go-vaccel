package vaccel

/*

#include <vaccel.h>

*/
import "C"

type Session struct {
	cSess C.struct_vaccel_session
}

func SessionInit(sess *Session, flags uint32) int {

	return int(C.vaccel_sess_init(&sess.cSess, C.uint32_t(flags))) //nolint:gocritic

}

func SessionRegister(sess *Session, res *Resource) int {

	return int(C.vaccel_sess_register(&sess.cSess, res.cRes)) //nolint:gocritic

}

func SessionFree(sess *Session) int {

	return int(C.vaccel_sess_free(&sess.cSess)) //nolint:gocritic

}
