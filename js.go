package cfscrape

import (
	"github.com/robertkrimen/otto"
)

var NodeExecutor = JSExecutorFunc(NodeExecute)
var vm = otto.New()

// A JSExecutor executes a javascript expression, and returns the resulting
// string. Note that the javascript that it must execute can be arbitrary.
type JSExecutor interface {
	ExecuteJS(javascript string) (string, error)
}

type JSExecutorFunc func(string) (string, error)

func (f JSExecutorFunc) ExecuteJS(s string) (string, error) {
	return f(s)
}

// Executes a javascript expression with node. This executor makes no security
// guarantees.
func NodeExecute(javascript string) (string, error) {
	result, err := vm.Run(javascript)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
