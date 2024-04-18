package parser

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/parser/types"
	"github.com/mitchellh/mapstructure"
)

type AST struct{ internal.AST }
type ASTs []*AST

func NewAST(val *internal.ASTResult) ASTs {
	ast := ASTs{}
	for _, a := range val.AST {
		ast = append(ast, &AST{a})
	}
	return ast
}

func (s ASTs) First() *AST {
	return s[0]
}

func (s ASTs) Len() int {
	return len(s)
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

func (s ASTs) FindFirst(t types.Statement) *AST {
	for _, ast := range s {
		if ast.Type() == t {
			return ast
		}
	}
	return nil
}

func (a AST) Type() types.Statement {
	return types.StatementFrom(a.M["type"].(string))
}

func (a AST) Is(s types.Statement) bool {
	return a.Type() == s
}

func (a AST) Use() (*types.Use, error) {
	if !a.Is(types.UseStatement) {
		return nil, fmt.Errorf("not a %s statement", types.UseStatement)
	}
	u := types.Use{}
	err := mapstructure.Decode(a.M, &u)
	return &u, err
}

func (a AST) Select() (*types.Select, error) {
	if !a.Is(types.SelectStatement) {
		return nil, fmt.Errorf("not a %s statement", types.SelectStatement)
	}
	s := types.Select{}
	err := mapstructure.Decode(a.M, &s)
	return &s, err
}

func (a AST) Insert() (*types.Insert, error) {
	if !a.Is(types.InsertStatement) {
		return nil, fmt.Errorf("not a %s statement", types.InsertStatement)
	}
	i := types.Insert{}
	err := mapstructure.Decode(a.M, &i)
	return &i, err
}

func (a AST) Update() (*types.Update, error) {
	if !a.Is(types.UpdateStatement) {
		return nil, fmt.Errorf("not a %s statement", types.UpdateStatement)
	}
	u := types.Update{}
	err := mapstructure.Decode(a.M, &u)
	return &u, err
}

func (a AST) Delete() (*types.Delete, error) {
	if !a.Is(types.DeleteStatement) {
		return nil, fmt.Errorf("not a %s statement", types.DeleteStatement)
	}
	d := types.Delete{}
	err := mapstructure.Decode(a.M, &d)
	return &d, err
}

func (a AST) Alter() (*types.Alter, error) {
	if !a.Is(types.AlterStatement) {
		return nil, fmt.Errorf("not a %s statement", types.AlterStatement)
	}
	al := types.Alter{}
	err := mapstructure.Decode(a.M, &al)
	return &al, err
}

func (a AST) Create() (*types.Create, error) {
	if !a.Is(types.CreateStatement) {
		return nil, fmt.Errorf("not a %s statement", types.CreateStatement)
	}
	c := types.Create{}
	err := mapstructure.Decode(a.M, &c)
	return &c, err
}

func (a AST) Drop() (*types.Drop, error) {
	if !a.Is(types.DropStatement) {
		return nil, fmt.Errorf("not a %s statement", types.DropStatement)
	}
	d := types.Drop{}
	err := mapstructure.Decode(a.M, &d)
	return &d, err
}
