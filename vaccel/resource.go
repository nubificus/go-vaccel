package vaccel


/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

type Resource struct {
	c_res *C.struct_vaccel_resource
}
