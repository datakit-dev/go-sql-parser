package internal

import (
	"github.com/dop251/goja"
)

type ASTResult struct {
	v   goja.Value
	AST []map[string]any
}

func NewASTResult(vm *goja.Runtime, val goja.Value) (*ASTResult, error) {
	var slice []map[string]any
	if _, ok := val.Export().([]any); ok {
		var ast []map[string]any
		err := vm.ExportTo(val, &ast)
		if err != nil {
			return nil, err
		}
		slice = append(slice, ast...)
	} else if _, ok := val.Export().(map[string]any); ok {
		var ast map[string]any
		err := vm.ExportTo(val, &ast)
		if err != nil {
			return nil, err
		}
		slice = append(slice, ast)
	}

	return &ASTResult{val, slice}, nil
}
