package testhelpers

import "testing"

func TestTypeName(t *testing.T) {
	cases := []struct {
		want string
		inst interface{}
	}{
		{"int32", int32(1)},
		{"int64", int64(1)},
		{"float64", float64(1)},
		{"string", ""},
		{"*testing.T", t},
		{"testing.T", *t},
	}

	for _, tcase := range cases {
		actual := TypeName(tcase.inst)
		expected := tcase.want
		if actual != expected {
			t.Error("TypeName()",
				NotEqualMsg(expected, actual))
		}
	}
}
