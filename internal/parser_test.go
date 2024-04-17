package internal

import (
	"fmt"
	"testing"

	"github.com/datakit-dev/go-sql-parser/internal/types"
)

func TestAllDatabases_Use(t *testing.T) {
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse(fmt.Sprintf("USE %s", "cool_db"), opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse("SELECT * FROM employees", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse("INSERT INTO employees (name, age) VALUES ('John Doe', 30)", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse("UPDATE employees SET age = 31 WHERE name = 'John Doe'", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse("DELETE FROM employees WHERE name = 'John Doe'", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		switch db {
		case BigQuery:
		case FlinkSQL:
		default:
			res, err := p.Parse("ALTER TABLE employees ADD COLUMN salary INT", opt)
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
				t.Errorf("%s AST length is greater than 1", db)
			}
			for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())

		var query string
		switch db {
		case Athena:
			query = "CREATE TABLE employees (id NUMERIC, name VARCHAR(255))"
		case BigQuery:
			query = `
				CREATE TABLE employees (
					id INT64,
					name TEXT,
					age INT64
				)
			`
		case Hive, TransactSQL, PostgreSQL, Redshift, MariaDB, MySQL, DB2, Trino:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case FlinkSQL:
			query = "CREATE TABLE employees (id BIGINT, name TEXT)"
		case Snowflake:
			query = "CREATE TABLE mytable (id INT, amount BIGINT);"
		case Sqlite:
			query = "CREATE TABLE employees (id INT, name TEXT)"
		default:
			query = "CREATE TABLE employees (id INT, name TEXT)"
		}

		if query != "" {
			res, err := p.Parse(query, opt)
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
				t.Errorf("%s AST length is greater than 1", db)
			}
			for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range Databases {
		opt := types.Option{}
		opt.SetDatabase(db.String())
		res, err := p.Parse("DROP TABLE employees", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.ASTResult != nil && len(res.ASTResult.ASTs) > 1 {
			t.Errorf("%s AST length is greater than 1", db)
		}
		for _, ast := range res.ASTResult.ASTs {
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
	p, err := NewParser()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	opt := types.Option{}
	opt.SetDatabase(BigQuery.String())

	res, err := p.Parse(`
		WITH t1 AS (
			SELECT a, b FROM schema1.table1
		), t2 AS (
			SELECT c, d FROM schema2.table2
		)
		SELECT a, b FROM t1
		UNION ALL
		SELECT c, d FROM t2
	`, opt)
	if err != nil {
		t.Error(fmt.Sprintf("%s failed", BigQuery), err)
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
	if res != nil && res.ASTResult == nil {
		t.Error("AST is nil")
	} else if len(res.ASTResult.ASTs) == 0 {
		t.Error("AST is empty")
	} else {
		for _, ast := range res.ASTResult.ASTs {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if !ast.Is(types.SelectStatement) {
					t.Error("AST type is not select:", ast.Type())
				} else {
					selectStmt, err := ast.Select()
					if err != nil {
						t.Error("Error decoding select statement:", err)
					}
					if selectStmt == nil {
						t.Error("Select statement is nil")
					}
				}
			}
		}
	}
}
