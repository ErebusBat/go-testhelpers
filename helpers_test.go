package testhelpers

import (
	"fmt"
	"testing"
)

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

// mISlice - Make Interface{} Slice is a helper function to aid in constructing test cases
func mISlice(a ...interface{}) []interface{} {
	return a
}

func TestExtractFormatAndArgs(t *testing.T) {
	cases := []struct {
		name   string
		ok     bool
		format string
		args   []interface{}
		subj   []interface{}
	}{
		{name: "Should not be OK on nil subj", ok: false, subj: nil},
		{"Zero args", true, "fmt", nil, mISlice("fmt")},
		{"One arg", true, "fmt", mISlice(1), mISlice("fmt", 1)},
		{"Two args", true, "fmt", mISlice(1, 2), mISlice("fmt", 1, 2)},
	}

	for idx, tcase := range cases {
		aOk, aFmt, _ := ExtractFormatAndArgs(tcase.subj...)
		msgBase := "Case #%d (%s), Variable %q Mismatch"
		if aOk != tcase.ok {
			t.Error(fmt.Sprintf(msgBase, idx, tcase.name, "OK"), NotEqualMsg(tcase.ok, aOk))
		}
		if aOk {
			if aFmt != tcase.format {
				t.Error(fmt.Sprintf(msgBase, idx, tcase.name, "Format"), NotEqualMsg(tcase.format, aFmt))
			}
		}
	}
}
