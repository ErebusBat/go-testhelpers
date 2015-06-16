package testhelpers

import "fmt"

type TestingT interface {
	Error(args ...interface{})
}

// Type that implementes testify.assert.TestingT to test if
// Errorf() was invoked
type TestingTRecorder struct {
	ErrorArgs      []interface{}
	ErrorWasCalled bool
}

func NewTestingTRecorder() TestingTRecorder {
	ttr := TestingTRecorder{}
	ttr.Reset()
	return ttr
}

func (ttr *TestingTRecorder) Error(args ...interface{}) {
	ttr.ErrorArgs = args
	ttr.ErrorWasCalled = true
}
func (ttr *TestingTRecorder) ErrorMessage() string {
	if !ttr.ErrorWasCalled {
		return ""
	}
	// Sprintln is what testing.T does, see:
	//   https://github.com/golang/go/blob/0266bc84940657b1e09f72bfe3d932f0344bc6a3/src/testing/testing.go#L359
	return fmt.Sprintln(ttr.ErrorArgs...)
}
func (ttr *TestingTRecorder) Reset() {
	ttr.ErrorArgs = nil
	ttr.ErrorWasCalled = false
}
