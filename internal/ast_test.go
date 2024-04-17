package internal

import (
	"testing"

	"github.com/datakit-dev/go-sql-parser/parser/types"
)

func TestAST_Type(t *testing.T) {
	tests := []struct {
		ast      AST
		expected types.Statement
	}{
		{
			ast:      AST{"type": "select"},
			expected: types.SelectStatement,
		},
		{
			ast:      AST{"type": "insert"},
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
		ast      AST
		stmt     types.Statement
		expected bool
	}{
		{
			ast:      AST{"type": "select"},
			stmt:     types.SelectStatement,
			expected: true,
		},
		{
			ast:      AST{"type": "insert"},
			stmt:     types.SelectStatement,
			expected: false,
		},
		{
			ast:      AST{"type": "foobar"},
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
				{"type": "select"},
				{"type": "insert"},
				{"type": "update"},
				{"type": "delete"},
				{"type": "create"},
				{"type": "drop"},
				{"type": "alter"},
				{"type": "select"},
				{"type": "foobar"},
				{"type": ""},
			},
			expected: ASTs([]AST{{"type": "select"}, {"type": "select"}}),
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
		expected AST
	}{
		{
			ast: ASTs{
				{"type": "select"},
				{"type": "insert"},
				{"type": "update"},
				{"type": "delete"},
				{"type": "create"},
				{"type": "drop"},
				{"type": "alter"},
				{"type": "foobar"},
				{"type": ""},
			},
			expected: AST{"type": "select"},
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
		ast      AST
		expected *types.Select
	}{
		{
			ast: AST{"type": "select"},
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
