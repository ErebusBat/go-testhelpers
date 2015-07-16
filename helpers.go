package testhelpers

import "fmt"

func TypeName(inst interface{}) string {
	return fmt.Sprintf("%T", inst)
}

// ExtractFormatAndArgs takes a variadic array and extracts the first
// argument as a string, leaving the rest in a []interface{}.
// The purpose of this is to extract Printf style parameters passed into
// functions as a ...interface{}.
// OK will return false if format and args can be passed to a Printf function
func ExtractFormatAndArgs(array ...interface{}) (ok bool, format string, args []interface{}) {
	l := len(array)
	if l == 0 {
		return false, format, args
	} else {
		format, ok = array[0].(string)
	}
	if !ok || l == 1 {
		// if not ok then there is nothing we can do without a format string
		// at this point with only one argument it was either a string or not
		return
	}

	args = array[1:]
	return
}

// Sprintf is a wrapper around ExtractFormatAndArgs and fmt.Sprintf
// This allows functions to accept Sprintf style format and args as a
// single variadic input and pass it to this Sprintf function.
// If the input array is empty or otherwise invalid then an empty string
// will be returned
func Sprintf(array ...interface{}) string {
	ok, f, a := ExtractFormatAndArgs(array...)
	if !ok {
		return ""
	}
	return fmt.Sprintf(f, a...)
}
