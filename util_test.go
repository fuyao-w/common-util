package common_util

import (
	"testing"
	"unsafe"
)

func TestLock(t *testing.T) {
	l := NewLockItem(1)
	t.Log(l.Get())
	func() {
		val := l.Lock()
		defer l.Unlock()
		t.Log(*val)
		*val = 2
	}()
	t.Log(l.Get())
}

func TestPtr(t *testing.T) {

	t.Log(*Ptr(1))
	t.Log(*Ptr("11"))
	t.Log(*Ptr(complex(1, 2)))
}

func TestZero(t *testing.T) {
	t.Log(Zero[string]())
	t.Log(Zero[struct {
		a int
	}]())
	t.Log(Zero[unsafe.Pointer]())
	t.Log(Zero[*unsafe.Pointer]())
}

func TestStr(t *testing.T) {
	t.Log(Str2Bytes("abcd"))
	t.Log(Bytes2Str([]byte("abcd")))
}

func TestMax(t *testing.T) {
	t.Log(Max(1, 2))
	t.Log(Max("1", "2"))
	t.Log(Max(2.9, 2.9))

	t.Log(Min(1, 2))
	t.Log(Min("1", "2"))
	t.Log(Min(2.9, 2.9))
	x := []int{4, 5, 1, 4, 6, 7, 6}
	SortSlice(x)
	t.Log(x)
	SortSlice(x, true)
	t.Log(x)
}
