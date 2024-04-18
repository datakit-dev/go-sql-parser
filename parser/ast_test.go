package parser

import (
	"testing"

	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/parser/types"
)

func newAST(m map[string]any) *AST {
	return &AST{internal.AST{M: m}}
}

func TestAST_Type(t *testing.T) {
	tests := []struct {
		ast      *AST
		expected types.Statement
	}{
		{
			ast:      newAST(map[string]any{"type": "select"}),
			expected: types.SelectStatement,
		},
		{
			ast:      newAST(map[string]any{"type": "insert"}),
			expected: types.InsertStatement,
		},
	}

	for _, test := range tests {
		result := test.ast.Type()
		if result != test.expected {
			t.Errorf("AST.Type returned %v, expected %v", result, test.expected)
		}
	}
}

func TestAST_Is(t *testing.T) {
	tests := []struct {
		ast      *AST
		stmt     types.Statement
		expected bool
	}{
		{
			ast:      newAST(map[string]any{"type": "select"}),
			stmt:     types.SelectStatement,
			expected: true,
		},
		{
			ast:      newAST(map[string]any{"type": "insert"}),
			stmt:     types.SelectStatement,
			expected: false,
		},
		{
			ast:      newAST(map[string]any{"type": "foobar"}),
			stmt:     types.SelectStatement,
			expected: false,
		},
	}

	for _, test := range tests {
		result := test.ast.Is(test.stmt)
		if result != test.expected {
			t.Errorf("AST.Is returned %v, expected %v", result, test.expected)
		}
	}
}

func TestASTSlice_FindAll(t *testing.T) {
	tests := []struct {
		ast      ASTs
		expected ASTs
	}{
		{
			ast: ASTs{
				newAST(map[string]any{"type": "select"}),
				newAST(map[string]any{"type": "insert"}),
				newAST(map[string]any{"type": "update"}),
				newAST(map[string]any{"type": "delete"}),
				newAST(map[string]any{"type": "create"}),
				newAST(map[string]any{"type": "drop"}),
				newAST(map[string]any{"type": "alter"}),
				newAST(map[string]any{"type": "select"}),
				newAST(map[string]any{"type": "foobar"}),
				newAST(map[string]any{"type": ""}),
			},
			expected: ASTs{
				newAST(map[string]any{"type": "select"}),
				newAST(map[string]any{"type": "select"}),
			},
		},
	}

	for _, test := range tests {
		result := test.ast.FindAll(types.SelectStatement)
		if len(result) != len(test.expected) {
			t.Errorf("ASTSlice.Find returned %v, expected %v", result, test.expected)
		}
	}
}

func TestASTSlice_FindFirst(t *testing.T) {
	tests := []struct {
		ast      ASTs
		expected *AST
	}{
		{
			ast: ASTs{
				newAST(map[string]any{"type": "select"}),
				newAST(map[string]any{"type": "insert"}),
				newAST(map[string]any{"type": "update"}),
				newAST(map[string]any{"type": "delete"}),
				newAST(map[string]any{"type": "create"}),
				newAST(map[string]any{"type": "drop"}),
				newAST(map[string]any{"type": "alter"}),
				newAST(map[string]any{"type": "foobar"}),
				newAST(map[string]any{"type": ""}),
			},
			expected: newAST(map[string]any{"type": "select"}),
		},
	}

	for _, test := range tests {
		result := test.ast.FindFirst(types.SelectStatement)
		if result.Type() != test.expected.Type() {
			t.Errorf("ASTSlice.Find returned %v, expected %v", result, test.expected)
		}
	}
}

func TestAST_Select(t *testing.T) {
	tests := []struct {
		ast      *AST
		expected *types.Select
	}{
		{
			ast: newAST(map[string]any{"type": "select"}),
			expected: &types.Select{
				Type: types.SelectStatement,
			},
		},
	}

	for _, test := range tests {
		result, err := test.ast.Select()
		if err != nil {
			t.Errorf("AST.Select returned error: %v", err)
		}
		if result.Type != test.expected.Type {
			t.Errorf("AST.Select returned %v, expected %v", result, test.expected)
		}
	}
}
