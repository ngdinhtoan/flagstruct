package flagstruct

import "testing"

func TestGetStructValue(t *testing.T) {
	type A struct {
		X string `flag:"x"`
		Y string `flag:"y"`
		Z string `flag:"z"`
	}

	type AA struct {
		XX string `flag:"xx"`
	}

	type B struct {
		A
		*AA
		M  string `flag:"m"`
		YY string `flag:"-"`
		n  string // unexport string
	}

	if fields := structFields(&B{}); len(fields) != 5 {
		t.Fatalf("len of field should be 5, got %d", len(fields))
	}
}

func TestNotStruct(t *testing.T) {
	defer func() {
		if err := recover(); err != "not struct" {
			t.Fail()
		}
	}()

	structVal("string")
}

func TestIsStructPointer(t *testing.T) {
	type A struct{}
	var a *A

	if isStructPointer(a) == true {
		t.Fail()
	}

	a = &A{}
	if isStructPointer(a) == false {
		t.Fail()
	}
}
