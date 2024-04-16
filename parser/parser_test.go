package parser_test

import (
	"fmt"
	"testing"

	"github.com/datakit-dev/go-sql-parser/parser"
	"github.com/datakit-dev/go-sql-parser/parser/types"
)

func TestAllDatabases_Use(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse(fmt.Sprintf("USE %s", "cool_db"), parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.UseStatement {
					t.Error("AST type is not use")
				}
			}
		}
	}
}

func TestAllDatabases_Select(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse("SELECT * FROM employees", parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.SelectStatement {
					t.Error("AST type is not select")
				}
			}
		}
	}
}

func TestAllDatabases_Insert(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse("INSERT INTO employees (name, age) VALUES ('John Doe', 30)", parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.InsertStatement {
					t.Error("AST type is not insert")
				}
			}
		}
	}
}

func TestAllDatabases_Update(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse("UPDATE employees SET age = 31 WHERE name = 'John Doe'", parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.UpdateStatement {
					t.Error("AST type is not update")
				}
			}
		}
	}
}

func TestAllDatabases_Delete(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse("DELETE FROM employees WHERE name = 'John Doe'", parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.DeleteStatement {
					t.Error("AST type is not delete")
				}
			}
		}
	}
}

func TestAllDatabases_Alter(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		switch db {
		case parser.DatabaseBigQuery:
		case parser.DatabaseFlinkSQL:
		default:
			res, err := p.Parse("ALTER TABLE employees ADD COLUMN salary INT", parser.WithDatabase(db))
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.AST != nil && len(res.AST) > 1 {
				t.Errorf("%s AST length is greater than 1", db)
			}
			for _, ast := range res.AST {
				if ast == nil {
					t.Error("AST is nil")
				} else {
					if ast.Type() != types.AlterStatement {
						t.Error("AST type is not alter")
					}
				}
			}
		}
	}
}

func TestAllDatabases_Create(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		var query string
		switch db {
		case parser.DatabaseBigQuery:
			query = `
				CREATE TABLE employees (
					id INT64,
					name STRING,
					age INT64
				)
			`
		case parser.DatabaseDB2:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case parser.DatabaseMariaDB, parser.DatabaseMySQL:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case parser.DatabaseNoql:
		case parser.DatabasePostgreSQL, parser.DatabaseRedshift:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case parser.DatabaseSnowflake:
			query = "CREATE TABLE mytable (amount NUMBER);"
		case parser.DatabaseSqlite:
			query = "CREATE TABLE employees (id INT, name TEXT)"
		case parser.DatabaseTransactSQL:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case parser.DatabaseTrino:
			query = "CREATE TABLE employees (id INT, name VARCHAR)"
		default:
			query = "CREATE TABLE employees (id INT, name STRING)"
		}

		if query != "" {
			res, err := p.Parse(query, parser.WithDatabase(db))
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.AST != nil && len(res.AST) > 1 {
				t.Errorf("%s AST length is greater than 1", db)
			}
			for _, ast := range res.AST {
				if ast == nil {
					t.Error("AST is nil")
				} else {
					if ast.Type() != types.CreateStatement {
						t.Error("AST type is not create")
					}
				}
			}
		}
	}
}

func TestAllDatabases_Drop(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		res, err := p.Parse("DROP TABLE employees", parser.WithDatabase(db))
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST != nil && len(res.AST) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.DropStatement {
					t.Error("AST type is not drop")
				}
			}
		}
	}
}

func TestBigQuery_With_Select(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	res, err := p.Parse(`
		WITH t1 AS (
			SELECT a, b FROM schema1.table1
		), t2 AS (
			SELECT c, d FROM schema2.table2
		)
		SELECT a, b FROM t1
		UNION ALL
		SELECT c, d FROM t2
	`, parser.WithDatabase(parser.DatabaseBigQuery))
	if err != nil {
		t.Error(fmt.Sprintf("%s failed", parser.DatabaseBigQuery), err)
	}

	if res == nil {
		t.Error("Result is nil")
	}
	if res != nil && res.ColumnList == nil {
		t.Error("ColumnList is nil")
	}
	if res != nil && len(res.ColumnList) == 0 {
		t.Error("ColumnList is empty")
	}
	if res != nil && res.TableList == nil {
		t.Error("TableList is nil")
	}
	if res != nil && res.TableList != nil && len(res.TableList) == 0 {
		t.Error("TableList is empty")
	}
	if res != nil && res.AST == nil {
		t.Error("AST is nil")
	} else if len(res.AST) == 0 {
		t.Error("AST is empty")
	} else {
		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if ast.Type() != types.SelectStatement {
					t.Error("AST type is not select")
				}
			}
		}
	}
}
