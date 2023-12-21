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