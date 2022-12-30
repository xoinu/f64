// Package f64 provides utilities to encode/decode integers and bitset into/from float64
package f64

import (
	"unsafe"

	"github.com/xoinu/bitset"
)

type Float64 float64

func (f *Float64) SetInt64(i int64) {
	*(*int64)(unsafe.Pointer(f)) = i
}

func (f *Float64) Int64() int64 {
	return *(*int64)(unsafe.Pointer(f))
}

func (f *Float64) SetUint64(i uint64) {
	*(*uint64)(unsafe.Pointer(f)) = i
}

func (f *Float64) Uint64() uint64 {
	return *(*uint64)(unsafe.Pointer(f))
}

// SetInt32 sets int32 integers to float64
func (f *Float64) SetInt32(high int32, low int32) {
	p := uintptr(unsafe.Pointer(f))
	*(*int32)(unsafe.Pointer(p + 0)) = low
	*(*int32)(unsafe.Pointer(p + 4)) = high
}

// SetInt32 sets an int32 integer to the high bytes of float64
func (f *Float64) SetInt32H(high int32) {
	p := uintptr(unsafe.Pointer(f))
	*(*int32)(unsafe.Pointer(p + 4)) = high
}

// SetInt32 sets an int32 integer to the low bytes of float64
func (f *Float64) SetInt32L(low int32) {
	p := uintptr(unsafe.Pointer(f))
	*(*int32)(unsafe.Pointer(p + 0)) = low
}

// Int32 returns int32 integers encoded in float64
func (f Float64) Int32() (high int32, low int32) {
	p := uintptr(unsafe.Pointer(&f))
	low = *(*int32)(unsafe.Pointer(p + 0))
	high = *(*int32)(unsafe.Pointer(p + 4))
	return
}

// Int32L returns int32 integers encoded in the high bytes of float64
func (f Float64) Int32H() int32 {
	p := uintptr(unsafe.Pointer(&f))
	return *(*int32)(unsafe.Pointer(p + 4))
}

// Int32L returns int32 integers encoded in the low bytes of float64
func (f Float64) Int32L() int32 {
	p := uintptr(unsafe.Pointer(&f))
	return *(*int32)(unsafe.Pointer(p + 0))
}

// SetUint32 sets uint32 integers to float64
func (f *Float64) SetUint32(high uint32, low uint32) {
	p := uintptr(unsafe.Pointer(f))
	*(*uint32)(unsafe.Pointer(p + 0)) = low
	*(*uint32)(unsafe.Pointer(p + 4)) = high
}

// SetUint32 sets an uint32 integer to the high bytes of float64
func (f *Float64) SetUint32H(high uint32) {
	p := uintptr(unsafe.Pointer(f))
	*(*uint32)(unsafe.Pointer(p + 4)) = high
}

// SetUint32 sets an uint32 integer to the low bytes of float64
func (f *Float64) SetUint32L(low uint32) {
	p := uintptr(unsafe.Pointer(f))
	*(*uint32)(unsafe.Pointer(p + 0)) = low
}

// Uint32 returns uint32 integers encoded in float64
func (f Float64) Uint32() (high uint32, low uint32) {
	p := uintptr(unsafe.Pointer(&f))
	low = *(*uint32)(unsafe.Pointer(p + 0))
	high = *(*uint32)(unsafe.Pointer(p + 4))
	return
}

// Uint32L returns uint32 integers encoded in the high bytes of float64
func (f Float64) Uint32H() uint32 {
	p := uintptr(unsafe.Pointer(&f))
	return *(*uint32)(unsafe.Pointer(p + 4))
}

// Uint32L returns uint32 integers encoded in the low bytes of float64
func (f Float64) Uint32L() uint32 {
	p := uintptr(unsafe.Pointer(&f))
	return *(*uint32)(unsafe.Pointer(p + 0))
}

// SetBit sets an bool to float64. n must satisfy 0 <= n < 64.
func (f *Float64) SetBit(n int, b bool) {
	(*bitset.BitSet64)(unsafe.Pointer(f)).Set(n, b)
}

// SetBit sets an bool to the high bytes of float64. n must satisfy 0 <= n < 32.
func (f *Float64) SetBitH(n int, b bool) {
	p := uintptr(unsafe.Pointer(f))
	(*bitset.BitSet)(unsafe.Pointer(p+4)).Set(n, b)
}

// SetBit sets an bool integer to the low bytes of float64. n must satisfy 0 <= n < 32. Ignored otherwise.
func (f *Float64) SetBitL(n int, b bool) {
	p := uintptr(unsafe.Pointer(f))
	(*bitset.BitSet)(unsafe.Pointer(p)).Set(n, b)
}

// TestBit returns bool encoded in float64. n must satisfy 0 <= n < 64. Ignored otherwise.
func (f Float64) TestBit(n int) bool {
	return (*bitset.BitSet64)(unsafe.Pointer(&f)).Test(n)
}

// TestBitH returns bool encoded in the high bytes of float64. n must satisfy 0 <= n < 32. Ignored otherwise.
func (f Float64) TestBitH(n int) bool {
	p := uintptr(unsafe.Pointer(&f))
	return (*bitset.BitSet)(unsafe.Pointer(p + 4)).Test(n)
}

// TestBitL returns bool encoded in the low bytes of float64. n must satisfy 0 <= n < 32. Ignored otherwise.
func (f Float64) TestBitL(n int) bool {
	p := uintptr(unsafe.Pointer(&f))
	return (*bitset.BitSet)(unsafe.Pointer(p + 0)).Test(n)
}

// SetInt8 encodes int8 into float64 as if it were [8]int8. n must satisfy 0 <= n < 7. Ignored otherwise.
func (f *Float64) SetInt8(n int, i int8) {
	if 0 <= n && n < 8 {
		p := uintptr(unsafe.Pointer(f))
		*(*int8)(unsafe.Pointer(p + uintptr(n))) = i
	}
}

// Int8 decodes int8 from float64 as if it were [8]int8. n must satisfy 0 <= n < 7. Ignored otherwise.
func (f *Float64) Int8(n int) int8 {
	if 0 <= n && n < 8 {
		p := uintptr(unsafe.Pointer(f))
		return *(*int8)(unsafe.Pointer(p + uintptr(n)))
	}
	return 0
}

// SetUint8 encodes int8 into float64 as if it were [8]uint8. n must satisfy 0 <= n < 7. Ignored otherwise.
func (f *Float64) SetUint8(n int, i uint8) {
	if 0 <= n && n < 8 {
		p := uintptr(unsafe.Pointer(f))
		*(*uint8)(unsafe.Pointer(p + uintptr(n))) = i
	}
}

// Uint8 decodes uint8 from float64 as if it were [8]uint8. n must satisfy 0 <= n < 7. Ignored otherwise.
func (f *Float64) Uint8(n int) uint8 {
	if 0 <= n && n < 8 {
		p := uintptr(unsafe.Pointer(f))
		return *(*uint8)(unsafe.Pointer(p + uintptr(n)))
	}
	return 0
}
