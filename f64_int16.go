package f64

import "unsafe"

// SetInt16 encodes int16 into float64 as if it were [4]int16. n must satisfy 0 <= n < 4. Ignored otherwise.
func (f *Float64) SetInt16(n int, i int16) {
	if 0 <= n && n < 4 {
		p := uintptr(unsafe.Pointer(f))
		*(*int16)(unsafe.Pointer(p + uintptr(n*2))) = i
	}
}

// Int16 decodes int16 from float64 as if it were [8]int16. n must satisfy 0 <= n < 4. Ignored otherwise.
func (f *Float64) Int16(n int) int16 {
	if 0 <= n && n < 4 {
		p := uintptr(unsafe.Pointer(f))
		return *(*int16)(unsafe.Pointer(p + uintptr(n*2)))
	}
	return 0
}

// SetUint16 encodes uint16 into float64 as if it were [4]uint16. n must satisfy 0 <= n < 4. Ignored otherwise.
func (f *Float64) SetUint16(n int, i uint16) {
	if 0 <= n && n < 4 {
		p := uintptr(unsafe.Pointer(f))
		*(*uint16)(unsafe.Pointer(p + uintptr(n*2))) = i
	}
}

// Uint16 decodes uint16 from float64 as if it were [8]uint16. n must satisfy 0 <= n < 4. Ignored otherwise.
func (f *Float64) Uint16(n int) uint16 {
	if 0 <= n && n < 4 {
		p := uintptr(unsafe.Pointer(f))
		return *(*uint16)(unsafe.Pointer(p + uintptr(n*2)))
	}
	return 0
}
