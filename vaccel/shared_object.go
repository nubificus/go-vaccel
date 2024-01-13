package vaccel

/*

#include <vaccel.h>

*/
import "C"

type SharedObject struct {
	cObj C.struct_vaccel_shared_object
}

func SharedObjectNew(obj *SharedObject, path string) int {

	return int(C.vaccel_shared_object_new(&obj.cObj, C.CString(path))) //nolint:gocritic
}

func (obj *SharedObject) GetResource() *Resource {

	res := new(Resource)
	res.cRes = obj.cObj.resource

	return res

}
