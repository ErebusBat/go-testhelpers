package examples

import (
	"testing"

	th "github.com/ErebusBat/go-testhelpers"
)

func ExampleNotEqual(t *testing.T) {
	expected := "The quick brown fox jumped over the lazy dog"
	actual := "The quick brown fox jumped over that lazy dog"
	if expected != actual {
		t.Error(th.NotEqualMsg(expected, actual))
	}

	// Additional information, and value
	if expected != actual {
		t.Error("Addition context",
			th.ES_VALUE, "Secret Value",
			th.NotEqualMsg(expected, actual))
	}
}
func ExampleFieldTest(t *testing.T) {
	cases := []struct {
		expected interface{}
		actual   interface{}
		msg      string
	}{
		{1234, 1234.0, "int => float"},
		{"Billy Bob Thorton", "Billy Bob THorton", "Last Name"},
		{true, false, "bool"},
	}
	for _, tc := range cases {
		if tc.expected != tc.actual {
			t.Error(tc.msg, th.NotEqualMsg(tc.expected, tc.actual))
		}
	}
}
