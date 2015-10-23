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
		M string `flag:"m"`
		n string // unexport string
	}

	if fields := structFields(&B{}); len(fields) != 4 {
		t.Fail()
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
