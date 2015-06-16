package testhelpers

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRawUsage(t *testing.T) {
	// Testing Table
	tt := []struct {
		name     string
		expected string
		args     []interface{}
	}{
		{"ES_EXPECTED", "\nexpected: TEST\n", []interface{}{ES_EXPECTED, "TEST"}},
		{"ES_GOT", "\n  actual: TEST\n", []interface{}{ES_GOT, "TEST"}},
		{"ES_ARGS", "\n    args: TEST\n", []interface{}{ES_ARGS, "TEST"}},
		{"ES_SQL", "\n     sql: TEST\n", []interface{}{ES_SQL, "TEST"}},
		{"ES_ERR", "\n     err: TEST\n", []interface{}{ES_ERR, "TEST"}},
		{"ES_VALUE", "\n   value: TEST\n", []interface{}{ES_VALUE, "TEST"}},
		{"ES_COUNT", "\n   count: TEST\n", []interface{}{ES_COUNT, "TEST"}},
	}

	ttr := NewTestingTRecorder()
	for _, tstCase := range tt {
		ttr.Error(tstCase.args...)
		actual := ttr.ErrorMessage()
		if actual != tstCase.expected {
			t.Error(tstCase.name, getNotMatchErrMessage(tstCase.expected, actual))
		}
	}
}

func TestNotEqualMsg(t *testing.T) {
	expected := `
expected: ONE
  actual: TWO`
	actual := NotEqualMsg("ONE", "TWO")
	if !areMessageEquivilent(expected, actual) {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

func TestTypeNotEqualMsg(t *testing.T) {
	var etype float64
	var atype int64
	expected := `
expected: float64
  actual: int64`
	actual := TypeNotEqualMsg(etype, atype)
	if !areMessageEquivilent(expected, actual) {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

func TestSqlArgsMsg(t *testing.T) {
	var args = []interface{}{
		1, 2.0, 1.1, true, "false",
	}
	expected := `
     sql: SELECT * FROM TABLE
    args: [1 2 1.1 true false]`
	actual := SqlArgsMsg("SELECT * FROM TABLE", args)
	if !areMessageEquivilent(expected, actual) {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

func TestNotEqualMsgWithRaw(t *testing.T) {
	ttr := NewTestingTRecorder()
	expected := `
   value: VALUE
expected: ONE
  actual: TWO`

	ttr.Error(ES_VALUE, "VALUE", NotEqualMsg("ONE", "TWO"))
	actual := ttr.ErrorMessage()
	if !areMessageEquivilent(expected, actual) {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

func TestUnexpectedErrMsg(t *testing.T) {
	expected := `Unexpected Error
     err: ERROR MESSAGE`

	actual := UnexpectedErrMsg("ERROR MESSAGE")
	if !areMessageEquivilent(expected, actual) {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

func TestKillTrailingWhitespaceHelper(t *testing.T) {
	expected := "The sly fox\nhe is sneaky"
	actual := killTrailingWhitespace("The sly fox \nhe is sneaky ")

	if expected != actual {
		t.Error(getNotMatchErrMessage(expected, actual))
	}
}

////////////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////////////

// helper to build message, given that we are testing builing a message
func getNotMatchErrMessage(expected, actual string) string {
	return fmt.Sprintf("Unexpected output (expected,actual):\n==>%s<==\n==>%s<==",
		expected,
		actual)
}

// Trailing whitespace does not matter this implementaion, Sprintln adds it
// and Sublime Text kills it on save, making it a PIA to track down
var reKillWS1 = regexp.MustCompile(`\s+\n`)
var reKillWS2 = regexp.MustCompile(`\s+$`)

func killTrailingWhitespace(v string) string {
	s := reKillWS1.ReplaceAllString(v, "\n")
	s = reKillWS2.ReplaceAllString(s, "")
	// fmt.Println(s)
	return s
}

func areMessageEquivilent(expected, actual string) bool {
	pass := (killTrailingWhitespace(expected) == killTrailingWhitespace(actual))
	return pass
}
