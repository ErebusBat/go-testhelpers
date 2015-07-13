package testhelpers

import "fmt"

const (
	// ERROR MESSAGEs - Used as first argument in Error() call
	EM_NEED_ERR = "Should have gotten an error, but didn't"
	EM_UN_ERR   = "Unexpected Error"

	// ERROR STRINGS - embed in
	// Could use a VAR block, fmt and %8s, but this is just as easy
	ES_EXPECTED = "\nexpected:"
	ES_GOT      = "\n  actual:"
	ES_ARGS     = "\n    args:"
	ES_SQL      = "\n     sql:"
	ES_ERR      = "\n     err:"
	ES_VALUE    = "\n   value:"
	ES_COUNT    = "\n   count:"
)

// Supply expected and actual values and a pretty formatted string will be returned that can be passed into t.Error()
func NotEqualMsg(expected, actual interface{}) string {
	return fmt.Sprintln(ES_EXPECTED, expected, ES_GOT, actual)
}

// Same as NotEqualMsg, except that the type names will be printed instead
func TypeNotEqualMsg(expected, actual interface{}) string {
	eType := TypeName(expected)
	aType := TypeName(actual)
	return fmt.Sprintln(ES_EXPECTED, eType, ES_GOT, aType)
}

func UnexpectedErrMsg(err string) string {
	return fmt.Sprintln(EM_UN_ERR, ES_ERR, err)
}
