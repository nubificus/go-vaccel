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

type Resource struct {
	c_res *C.struct_vaccel_resource
}