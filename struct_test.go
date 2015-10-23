package flagstruct

import "testing"

func TestGetStructValue(t *testing.T) {
	type A struct {
		X string `flag:"z"`
		Y string `flag:"y"`
		Z string `flag:"z"`
	}

	type B struct {
		A
		M string `flag:"m"`
	}

	if fields := structFields(&B{}); len(fields) != 4 {
		t.Fail()
	}
}
