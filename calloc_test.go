package calloc

import "testing"

import (
	"reflect"
)


func TestMalloc(t *testing.T) {
	var b []byte
	b = Make(reflect.TypeOf(b), 24, 24).([]byte)
	defer Free(b)
	if len(b) != 24 {
		t.Fatal("b was the wrong length")
	}
	if cap(b) != 24 {
		t.Fatal("b was the wrong capacity")
	}
}

type testStruct struct {
	a *int
	b int
	c float32
}

func TestMake(t *testing.T) {
	var s []testStruct
	s = Make(reflect.TypeOf(s), 0, 10).([]testStruct)
	defer Free(s)

	if s == nil {
		t.Fatal("s was nil")
	}
	if len(s) != 0 {
		t.Fatal("len was wrong")
	}
	if cap(s) != 10 {
		t.Fatal("cap was wrong")
	}
}

func TestNew(t *testing.T) {
	var s *testStruct
	s = New(reflect.TypeOf(s)).(*testStruct)
	Free(s)
}

