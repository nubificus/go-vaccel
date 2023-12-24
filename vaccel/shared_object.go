package vaccel


/*

#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"


type SharedObject struct
{
	c_obj C.struct_vaccel_shared_object
}

func SharedObjectNew(obj *SharedObject, path string) int {

	return int(C.vaccel_shared_object_new(&obj.c_obj, C.CString(path)))
}

func (obj *SharedObject) GetResource() *Resource {

	res := new(Resource)
	res.c_res = obj.c_obj.resource

	return res

}
