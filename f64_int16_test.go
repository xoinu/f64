package f64

import "testing"

func TestInt16(t *testing.T) {
	var f Float64
	f.SetInt16(0, -1)
	if i := f.Int16(0); i == -1 {
		//OK
	} else {
		t.Fatal("SetInt16/Int16 failed")
	}
	for i := 0; i < 4; i++ {
		f.SetInt16(i, -1)
	}
	if h, l := f.Int32(); h == -1 && l == -1 {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}

	if f.Int16(-1) != 0 {
		t.Fatal("invalid index must be ignored")
	}
}

func TestUint16(t *testing.T) {
	var f Float64
	f.SetUint16(0, 0xffff)
	if i := f.Uint16(0); i == 0xffff {
		//OK
	} else {
		t.Fatal("SetUint16/Uint16 failed")
	}
	for i := 0; i < 4; i++ {
		f.SetUint16(i, 0xffff)
	}
	if h, l := f.Uint32(); h == 0xffffffff && l == 0xffffffff {
		//OK
	} else {
		t.Fatalf("h == %d, l == %d", h, l)
	}

	if f.Uint16(0xffff) != 0 {
		t.Fatal("invalid index must be ignored")
	}
}
