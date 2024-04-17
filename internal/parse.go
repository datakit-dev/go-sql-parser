package internal

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/internal/types"
	"github.com/dop251/goja"
)

type ParseResult struct {
	*ASTResult
	types.TableColumnAst
}

func NewParseResult(vm *goja.Runtime, val goja.Value) (*ParseResult, error) {
	tca := types.TableColumnAst{}
	err := vm.ExportTo(val, &tca)
	if err != nil {
		return nil, err
	}

	rObj := val.ToObject(vm)
	if rObj == nil {
		return nil, fmt.Errorf("parse failed")
	}

	astVal := rObj.Get("ast")
	if astVal == nil {
		return nil, fmt.Errorf("parse failed")
	}

	ast, err := NewAST(vm, astVal)
	if err != nil {
		return nil, err
	}

	return &ParseResult{ast, tca}, nil
}
