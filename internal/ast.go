package internal

import (
	"fmt"

	"github.com/dop251/goja"
)

type AST struct {
	v goja.Value
	M map[string]any
}

type ASTResult struct {
	AST []AST
}

func NewASTResult(vm *goja.Runtime, val goja.Value) (*ASTResult, error) {
	var slice []AST
	if v, ok := val.Export().([]any); ok {
		for _, item := range v {
			if _, ok := item.(map[string]any); !ok {
				return nil, fmt.Errorf("invalid AST result: %v", item)
			}
			ast := item.(map[string]any)
			slice = append(slice, AST{vm.ToValue(ast), ast})
		}
	} else if _, ok := val.Export().(map[string]any); ok {
		var ast map[string]any
		err := vm.ExportTo(val, &ast)
		if err != nil {
			return nil, err
		}
		slice = append(slice, AST{val, ast})
	} else {
		return nil, fmt.Errorf("invalid AST result: %v", val.Export())
	}

	return &ASTResult{slice}, nil
}
