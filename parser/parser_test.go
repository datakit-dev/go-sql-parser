package parser_test

import (
	"testing"

	"github.com/datakit-dev/go-sql-parser/internal/types"
	"github.com/datakit-dev/go-sql-parser/parser"
)

func TestParse_Select(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	res, err := p.Parse("SELECT * FROM employees")
	if err != nil {
		t.Error("Failed to parse SQL", err)
	}

	if res != nil && res.First != nil && len(res.ASTs) > 1 {
		t.Error("AST length is greater than 1")
	}

	for _, ast := range res.ASTs {
		if ast == nil {
			t.Error("AST is nil")
		} else {
			if ast.Type() != types.SelectStatement {
				t.Error("AST type is not select")
			}
			if !ast.Is(types.SelectStatement) {
				t.Error("AST type is not select")
			}
		}
	}

	sel, err := res.First.Select()
	if err != nil {
		t.Error("Failed to get select statement", err)
	}

	if sel == nil {
		t.Error("Select statement is nil")
	} else if sel.Columns == nil {
		t.Error("Columns is nil")
	} else if len(sel.Columns) == 0 {
		t.Error("Columns is empty")
	}
}

func TestSqlify_Select(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	sql := "SELECT `name`, `title` FROM `employees`"
	res, err := p.Parse(sql)
	if err != nil {
		t.Error("Failed to parse SQL", err)
	}

	sql2, err := p.Sqlify(res.ASTResult)
	if err != nil {
		t.Error("Failed to astify SQL", err)
	}

	if sql2 != sql {
		t.Error("SQLs do not match\n", sql, "\n", sql2)
	}
}
