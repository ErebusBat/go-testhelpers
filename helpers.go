package testhelpers

import "fmt"

func TypeName(inst interface{}) string {
	return fmt.Sprintf("%T", inst)
}
