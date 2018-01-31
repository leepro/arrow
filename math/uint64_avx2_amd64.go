// Code generated by type_simd_amd64.go.tmpl.
// DO NOT EDIT.

// +build !noasm

package math

import (
	"unsafe"

	"github.com/influxdata/arrow/array"
)

//go:noescape
func _sum_uint64_avx2(buf, len, res unsafe.Pointer)

func sum_uint64_avx2(a *array.Uint64) uint64 {
	buf := a.Uint64Values()
	var (
		p1  = unsafe.Pointer(&buf[0])
		p2  = unsafe.Pointer(uintptr(len(buf)))
		res uint64
	)
	_sum_uint64_avx2(p1, p2, unsafe.Pointer(&res))
	return res
}
