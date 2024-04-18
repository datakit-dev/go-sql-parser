package parser

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/parser/types"
	"github.com/mitchellh/mapstructure"
)

type AST map[string]any
type ASTs []AST

type ASTResult struct {
	*internal.ASTResult
	AST ASTs
}

func NewASTResult(val *internal.ASTResult) *ASTResult {
	ast := ASTs{}
	for _, a := range val.AST {
		ast = append(ast, a)
	}
	return &ASTResult{val, ast}
}

func (s ASTs) First() AST {
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
