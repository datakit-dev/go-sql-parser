package internal

import (
	"fmt"

	"github.com/dop251/goja"
)

type TableColumnList struct {
	TableList  []string `json:"tableList"`
	ColumnList []string `json:"columnList"`
}

type ParseResult struct {
	*ASTResult
	TableColumnList
}

func NewParseResult(vm *goja.Runtime, val goja.Value) (*ParseResult, error) {
	res := ParseResult{}
	tcl := TableColumnList{}
	err := vm.ExportTo(val, &tcl)
	if err != nil {
		return nil, err
	}
	res.TableColumnList = tcl

	rObj := val.ToObject(vm)
	if rObj == nil {
		return nil, fmt.Errorf("parse failed")
	}

	astVal := rObj.Get("ast")
	if astVal == nil {
		return nil, fmt.Errorf("parse failed")
	}

	res.ASTResult, err = NewASTResult(vm, astVal)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
