package common_util

import (
	"reflect"
	"unsafe"
)

func Bytes2Str(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Bytes[T ~string](s T) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}
