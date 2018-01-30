// +build !noasm

package math

import (
	"unsafe"

	"github.com/influxdata/arrow/array"
)

//go:noescape
func _sum_float64_avx2(buf, len, res unsafe.Pointer)

func sum_float64_avx2(a *array.Float64) float64 {
	buf := a.Float64Values()
	var (
		p1  = unsafe.Pointer(&buf[0])
		p2  = unsafe.Pointer(uintptr(len(buf)))
		res float64
	)
	_sum_float64_avx2(p1, p2, unsafe.Pointer(&res))
	return res
}
