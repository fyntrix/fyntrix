package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

type Vips struct{}

func New() *Vips {
	return &Vips{}
}

func (*Vips) Version() string {
	return C.GoString(C.fyntrix_vips_version())
}
