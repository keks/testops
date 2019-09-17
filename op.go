package testops // import "github.com/keks/testops"

import "testing"

type Op interface {
	Do(*testing.T, interface{})
}

type DumpOp struct {
	Name string
	V    interface{}
}

func (op DumpOp) Do(t *testing.T, v interface{}) {
	t.Logf("%s: %#v", op.Name, op.V)
}
