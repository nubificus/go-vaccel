package vaccel

/*
#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

type Session struct {
	cSess C.struct_vaccel_session
}

func SessionInit(sess *Session, flags uint32) int {
	return int(C.vaccel_session_init(&sess.cSess, C.uint32_t(flags))) //nolint:gocritic
}

func SessionRelease(sess *Session) int {
	return int(C.vaccel_session_release(&sess.cSess)) //nolint:gocritic
}
