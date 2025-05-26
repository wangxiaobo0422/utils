package typeconv

import (
	"reflect"
	"unsafe"
)

// UnsafeSliceToString: zero-copy slice convert to string
func UnsafeSliceToString(b []byte) string {
	return *((*string)(unsafe.Pointer(&b)))
}

// UnsafeStringToSlice: zero-copy string cover to slice
func UnsafeStringToSlice(s string) (b []byte) {
	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Data = uintptr(p)
	hdr.Len = len(s)
	hdr.Cap = len(s)

	return b
}
