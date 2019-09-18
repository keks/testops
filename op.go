package testops // import "github.com/keks/testops"

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Op interface {
	Do(*testing.T, interface{})
}

type TestCase struct {
	Name string
	Ops  []Op
}

func (tc TestCase) Runner(v interface{}) func(*testing.T) {
	return func(t *testing.T) {
		for _, op := range tc.Ops {
			t.Logf("start %T", op)
			op.Do(t, v)
			t.Logf("done %T", op)
		}
	}
}

type Env struct {
	Name string
	Func func(TestCase) (func(*testing.T), error)
}

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

type DumpOp struct {
	Name string
	V    interface{}
}

func (op DumpOp) Do(t *testing.T, v interface{}) {
	t.Logf("%s: %#v", op.Name, op.V)
}
