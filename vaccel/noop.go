package vaccel

/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

func NoOp(sess *Session) int {

	csess := sess.cSess
	c_ret := C.vaccel_noop(&csess)

	return int(c_ret)

}
