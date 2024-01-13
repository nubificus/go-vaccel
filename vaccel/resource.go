package vaccel

/*

#include <vaccel.h>

*/
import "C"

type Resource struct {
	cRes *C.struct_vaccel_resource
}
