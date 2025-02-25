package vaccel

/*
#cgo pkg-config: vaccel
#cgo LDFLAGS: -lvaccel -ldl
#include <vaccel.h>

*/
import "C"

type ResourceType int

const (
	ResourceLib ResourceType = iota
	ResourceData
	ResourceModel
)

type Resource struct {
	cRes C.struct_vaccel_resource
}

func (t ResourceType) ToCEnum() C.vaccel_resource_type_t {
	return C.vaccel_resource_type_t(t)
}

func ResourceInit(res *Resource, path string, resType ResourceType) int {
	return int(C.vaccel_resource_init(&res.cRes, C.CString(path), resType.ToCEnum()))
}

func ResourceRegister(res *Resource, sess *Session) int {
	return int(C.vaccel_resource_register(&res.cRes, &sess.cSess)) //nolint:gocritic
}

func ResourceUnregister(res *Resource, sess *Session) int {
	return int(C.vaccel_resource_unregister(&res.cRes, &sess.cSess)) //nolint:gocritic
}

func ResourceRelease(res *Resource) int {
	return int(C.vaccel_resource_release(&res.cRes)) //nolint:gocritic
}
