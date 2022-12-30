package f64

import (
	"testing"
)

func TestInt(t *testing.T) {
	var f Float64
	f.SetInt64(123)
	if i := f.Int64(); i == 123 {
		//OK
	} else {
		t.Fatal("SetInt64/Int64 failed")
	}

	if h, l := f.Int32(); h == 0 && l == 123 {
		//OK
	} else {
		t.Fatal("SetInt64 failed")
	}

	f.SetInt32L(234)
	f.SetInt32H(456)
	if h, l := f.Int32(); h == 456 && l == 234 {
		//OK
	} else {
		t.Fatal("SetIntH32 failed")
	}
	if h := f.Int32H(); h == 456 {
		//OK
	} else {
		t.Fatal("Int32H failed")
	}
	if l := f.Int32L(); l == 234 {
		//OK
	} else {
		t.Fatal("Int32L failed")
	}

	f.SetInt64(0)
	if i := f.Int64(); i == 0 {
		//OK
	} else {
		t.Fatal("SetInt64/Int64 failed")
	}

	f.SetInt32(456, 234)
	if h, l := f.Int32(); h == 456 && l == 234 {
		//OK
	} else {
		t.Fatal("SetIntH32 failed")
	}
	if h := f.Int32H(); h == 456 {
		//OK
	} else {
		t.Fatal("Int32H failed")
	}
	if l := f.Int32L(); l == 234 {
		//OK
	} else {
		t.Fatal("Int32L failed")
	}
}

func TestUint(t *testing.T) {
	var f Float64
	f.SetUint64(123)
	if i := f.Uint64(); i == 123 {
		//OK
	} else {
		t.Fatal("SetUint64/Uint64 failed")
	}

	if h, l := f.Uint32(); h == 0 && l == 123 {
		//OK
	} else {
		t.Fatal("SetUint64 failed")
	}

	f.SetUint32L(234)
	f.SetUint32H(456)
	if h, l := f.Uint32(); h == 456 && l == 234 {
		//OK
	} else {
		t.Fatal("SetUintH32 failed")
	}
	if h := f.Uint32H(); h == 456 {
		//OK
	} else {
		t.Fatal("Uint32H failed")
	}
	if l := f.Uint32L(); l == 234 {
		//OK
	} else {
		t.Fatal("Uint32L failed")
	}

	f.SetUint64(0)
	if i := f.Uint64(); i == 0 {
		//OK
	} else {
		t.Fatal("SetUint64/Uint64 failed")
	}

	f.SetUint32(456, 234)
	if h, l := f.Uint32(); h == 456 && l == 234 {
		//OK
	} else {
		t.Fatal("SetUintH32 failed")
	}
	if h := f.Uint32H(); h == 456 {
		//OK
	} else {
		t.Fatal("Uint32H failed")
	}
	if l := f.Uint32L(); l == 234 {
		//OK
	} else {
		t.Fatal("Uint32L failed")
	}
}

func TestInt8(t *testing.T) {
	var f Float64
	f.SetInt8(0, -1)
	if i := f.Int8(0); i == -1 {
		//OK
	} else {
		t.Fatal("SetInt8/Int8 failed")
	}
	for i := 0; i < 4; i++ {
		f.SetInt8(i, -1)
	}
	if h, l := f.Int32(); h == 0 && l == -1 {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}
	for i := 4; i < 8; i++ {
		f.SetInt8(i, -1)
	}
	if h, l := f.Int32(); h == -1 && l == -1 {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}

	if f.Int8(-1) != 0 {
		t.Fatal("invalid index must be ignored")
	}
}

func TestUint8(t *testing.T) {
	var f Float64
	f.SetUint8(0, 0xff)
	if i := f.Uint8(0); i == 0xff {
		//OK
	} else {
		t.Fatal("SetUint8/Uint8 failed")
	}
	for i := 0; i < 4; i++ {
		f.SetUint8(i, 0xff)
	}
	if h, l := f.Uint32(); h == 0 && l == 0xffffffff {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}
	for i := 4; i < 8; i++ {
		f.SetUint8(i, 0xff)
	}
	if h, l := f.Uint32(); h == 0xffffffff && l == 0xffffffff {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}

	if f.Uint8(128) != 0 {
		t.Fatal("invalid index must be ignored")
	}
}

func TestBit(t *testing.T) {
	var f Float64
	for i := 0; i < 64; i++ {
		if f.TestBit(i) == true {
			t.Fatalf("f.TestBit(%d) must be false", i)
		}
	}
	f.SetBit(0, true)
	if f.TestBit(0) == false {
		t.Fatal("must be true!")
	}
	for i := 1; i < 64; i++ {
		f.SetBit(i, true)
	}
	if f.Int64() != -1 {
		t.Fatal("all bits must be flagged")
	}

	for i := 0; i < 32; i++ {
		f = 0
		f.SetBitH(i, true)
		if f.TestBitH(i) == false || f.TestBitL(i) == true {
			t.Fatal("SetBitH failed")
		}
		f.SetBitL(i, true)
		if f.TestBitH(i) == false || f.TestBitL(i) == false {
			t.Fatal("SetBitL failed")
		}
	}

	f = 0
	f.SetBit(123, true)
	if f.Int64() != 0 {
		t.Fatal("invalid index must be ignored")
	}
}
