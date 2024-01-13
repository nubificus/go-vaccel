package vaccel

/*

#include <vaccel.h>

*/
import "C"

func NoOp(sess *Session) int {

	csess := sess.cSess
	cRet := C.vaccel_noop(&csess) //nolint:gocritic

	return int(cRet)

}
