package testhelpers

import (
	"fmt"
	"strings"
	"testing"
)

type FieldTest struct {
	Name string
	Pass bool
	Eval interface{}
	Aval interface{}
}

func TestFields(t *testing.T, fieldNameFormat string, fields []FieldTest) bool {
	pass := true

	// Make sure that they passed in a %s, otherwise they won't
	// know which field failed
	if !strings.Contains(fieldNameFormat, "%s") {
		fieldNameFormat += "%s"
	}

	// Check the fields
	for _, field := range fields {
		if !field.Pass {
			pass = false
			errMessage := fmt.Sprintf(fieldNameFormat, field.Name)
			t.Error(errMessage, NotEqualMsg(field.Eval, field.Aval))
		}
	}

	return pass
}
