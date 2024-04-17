package internal

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/parser/types"
	"github.com/dop251/goja"
	"github.com/mitchellh/mapstructure"
)

type AST map[string]any
type ASTs []AST

type ASTResult struct {
	ASTs
	First AST
	v     goja.Value
}

func NewASTResult(vm *goja.Runtime, val goja.Value) (*ASTResult, error) {
	var first AST
	var slice ASTs

	if _, ok := val.Export().([]any); ok {
		var a []map[string]any
		err := vm.ExportTo(val, &a)
		if err != nil {
			return nil, err
		}
		for idx, ast := range a {
			if idx == 0 {
				first = AST(ast)
			}
			slice = append(slice, AST(ast))
		}
	} else if _, ok := val.Export().(map[string]any); ok {
		var a map[string]any
		err := vm.ExportTo(val, &a)
		if err != nil {
			return nil, err
		}
		first = AST(a)
		slice = append(slice, first)
	}

	return &ASTResult{slice, first, val}, nil
}

func (s ASTs) FindAll(t types.Statement) ASTs {
	var result ASTs
	for _, ast := range s {
		if ast.Type() == t {
			result = append(result, ast)
		}
	}
	return result
}

func (s ASTs) FindFirst(t types.Statement) AST {
	for _, ast := range s {
		if ast.Type() == t {
			return ast
		}
	}
	return nil
}

func (a AST) Type() types.Statement {
	return types.StatementFrom(a["type"].(string))
}

func (a AST) Is(s types.Statement) bool {
	return a.Type() == s
}

func (a AST) Use() (*types.Use, error) {
	if !a.Is(types.UseStatement) {
		return nil, fmt.Errorf("not a %s statement", types.UseStatement)
	}
	u := types.Use{}
	err := mapstructure.Decode(a, &u)
	return &u, err
}

func (a AST) Select() (*types.Select, error) {
	if !a.Is(types.SelectStatement) {
		return nil, fmt.Errorf("not a %s statement", types.SelectStatement)
	}
	s := types.Select{}
	err := mapstructure.Decode(a, &s)
	return &s, err
}

func (a AST) Insert() (*types.Insert, error) {
	if !a.Is(types.InsertStatement) {
		return nil, fmt.Errorf("not a %s statement", types.InsertStatement)
	}
	i := types.Insert{}
	err := mapstructure.Decode(a, &i)
	return &i, err
}

func (a AST) Update() (*types.Update, error) {
	if !a.Is(types.UpdateStatement) {
		return nil, fmt.Errorf("not a %s statement", types.UpdateStatement)
	}
	u := types.Update{}
	err := mapstructure.Decode(a, &u)
	return &u, err
}

func (a AST) Delete() (*types.Delete, error) {
	if !a.Is(types.DeleteStatement) {
		return nil, fmt.Errorf("not a %s statement", types.DeleteStatement)
	}
	d := types.Delete{}
	err := mapstructure.Decode(a, &d)
	return &d, err
}

func (a AST) Alter() (*types.Alter, error) {
	if !a.Is(types.AlterStatement) {
		return nil, fmt.Errorf("not a %s statement", types.AlterStatement)
	}
	al := types.Alter{}
	err := mapstructure.Decode(a, &al)
	return &al, err
}

func (a AST) Create() (*types.Create, error) {
	if !a.Is(types.CreateStatement) {
		return nil, fmt.Errorf("not a %s statement", types.CreateStatement)
	}
	c := types.Create{}
	err := mapstructure.Decode(a, &c)
	return &c, err
}

func (a AST) Drop() (*types.Drop, error) {
	if !a.Is(types.DropStatement) {
		return nil, fmt.Errorf("not a %s statement", types.DropStatement)
	}
	d := types.Drop{}
	err := mapstructure.Decode(a, &d)
	return &d, err
}
