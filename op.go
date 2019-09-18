package testops // import "github.com/keks/testops"

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Op is a single operation.
type Op interface {
	Do(*testing.T, interface{})
}

// TestCase is a named series of operations.
type TestCase struct {
	Name string
	Ops  []Op
}

// Runner returns a function that performs the operations in order.
// The parameter is passed in as state.
func (tc TestCase) Runner(v interface{}) func(*testing.T) {
	return func(t *testing.T) {
		for _, op := range tc.Ops {
			t.Logf("start %T", op)
			op.Do(t, v)
			t.Logf("done %T", op)
		}
	}
}

// Env is the environment an operation is executed in.
// It prepares and tears down the state of the test case.
type Env struct {
	Name string
	Func func(TestCase) (func(*testing.T), error)
}

// Run runs all combinations of environments and test cases.
func Run(t *testing.T, envs []Env, tcs []TestCase) {
	for _, env := range envs {
		t.Run(env.Name, func(t_ *testing.T) {
			for _, tc := range tcs {
				testFunc, err := env.Func(tc)
				require.NoError(t_, err, "error initiating test environment")

				t_.Run(tc.Name, testFunc)
			}
		})
	}
}

// DumpOp is a simple operation that logs a value.
type DumpOp struct {
	Name string
	V    interface{}
}

// Do is the interface function of Op.
func (op DumpOp) Do(t *testing.T, v interface{}) {
	t.Logf("%s: %#v", op.Name, op.V)
}
