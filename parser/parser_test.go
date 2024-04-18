package parser_test

import (
	"fmt"
	"testing"

	"github.com/datakit-dev/go-sql-parser/parser"
	"github.com/datakit-dev/go-sql-parser/parser/types"
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

	if res != nil && res.AST.Len() > 1 {
		t.Error("AST length is greater than 1")
	}

	for _, ast := range res.AST {
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

	sel, err := res.AST.First().Select()
	if err != nil {
		t.Error("Failed to get select statement", err)
	}

	if sel == nil {
		t.Error("Select statement is nil")
	} else if sel.Columns == nil {
		t.Error("Columns is nil")
	} else if len(sel.Columns) == 0 {
		t.Error("Columns is empty")
	} else if sel.From == nil {
		t.Error("From is nil")
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

	sql2, err := p.Sqlify(res.AST[0])
	if err != nil {
		t.Error("Failed to astify SQL", err)
	}

	if sql2 != sql {
		t.Error("SQLs do not match\n", sql, "\n", sql2)
	}
}

func TestAllDatabases_Use(t *testing.T) {
	p, err := parser.New()
	if err != nil {
		t.Errorf("Error loading parser: %v", err)
	}

	for _, db := range parser.Databases {
		opt := parser.WithDatabase(db)
		res, err := p.Parse(fmt.Sprintf("USE %s", "cool_db"), opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error("AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if !ast.Is(types.UseStatement) {
					t.Error("AST type is not use")
				} else {
					useStmt, err := ast.Use()
					if err != nil {
						t.Error("Error decoding use statement:", err)
					}
					if useStmt == nil {
						t.Error("Use statement is nil")
					}
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
		opt := parser.WithDatabase(db)
		res, err := p.Parse("SELECT * FROM employees", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error(fmt.Sprintf("%s failed", db), "AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error(fmt.Sprintf("%s failed", db), "AST is nil")
			} else {
				if !ast.Is(types.SelectStatement) {
					t.Error(fmt.Sprintf("%s failed", db), "AST type is not select")
				} else {
					selectStmt, err := ast.Select()
					if err != nil {
						t.Error(fmt.Sprintf("%s failed", db), "Error decoding select statement:", err)
					}
					if selectStmt == nil {
						t.Error(fmt.Sprintf("%s failed", db), "Select statement is nil")
					}
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
		opt := parser.WithDatabase(db)
		res, err := p.Parse("INSERT INTO employees (name, age) VALUES ('John Doe', 30)", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error(fmt.Sprintf("%s failed", db), "AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error(fmt.Sprintf("%s failed", db), "AST is nil")
			} else {
				if !ast.Is(types.InsertStatement) {
					t.Error(fmt.Sprintf("%s failed", db), "AST type is not insert")
				} else {
					insertStmt, err := ast.Insert()
					if err != nil {
						t.Error(fmt.Sprintf("%s failed", db), "Error decoding insert statement:", err)
					}
					if insertStmt == nil {
						t.Error(fmt.Sprintf("%s failed", db), "Insert statement is nil")
					}
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
		opt := parser.WithDatabase(db)
		res, err := p.Parse("UPDATE employees SET age = 31 WHERE name = 'John Doe'", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error("AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if !ast.Is(types.UpdateStatement) {
					t.Error("AST type is not update")
				} else {
					updateStmt, err := ast.Update()
					if err != nil {
						t.Error("Error decoding update statement:", err)
					}
					if updateStmt == nil {
						t.Error("Update statement is nil")
					}
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
		opt := parser.WithDatabase(db)
		res, err := p.Parse("DELETE FROM employees WHERE name = 'John Doe'", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error("AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if !ast.Is(types.DeleteStatement) {
					t.Error("AST type is not delete")
				} else {
					deleteStmt, err := ast.Delete()
					if err != nil {
						t.Error("Error decoding delete statement:", err)
					}
					if deleteStmt == nil {
						t.Error("Delete statement is nil")
					}
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
		opt := parser.WithDatabase(db)
		switch db {
		case parser.BigQuery:
		case parser.FlinkSQL:
		default:
			res, err := p.Parse("ALTER TABLE employees ADD COLUMN salary INT", opt)
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.AST.Len() > 1 {
				t.Error("AST length is greater than 1")
			}

			for _, ast := range res.AST {
				if ast == nil {
					t.Error("AST is nil")
				} else {
					if !ast.Is(types.AlterStatement) {
						t.Error("AST type is not alter")
					} else {
						alterStmt, err := ast.Alter()
						if err != nil {
							t.Error("Error decoding alter statement:", err)
						}
						if alterStmt == nil {
							t.Error("Alter statement is nil")
						}
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
		opt := parser.WithDatabase(db)

		var query string
		switch db {
		case parser.Athena:
			query = "CREATE TABLE employees (id NUMERIC, name VARCHAR(255))"
		case parser.BigQuery:
			query = `
				CREATE TABLE employees (
					id INT64,
					name STRING,
					age INT64
				)
			`
		case parser.Hive, parser.TransactSQL, parser.PostgreSQL, parser.Redshift, parser.MariaDB, parser.MySQL, parser.DB2, parser.Trino:
			query = "CREATE TABLE employees (id INT, name VARCHAR(255))"
		case parser.FlinkSQL:
			query = "CREATE TABLE employees (id BIGINT, name TEXT)"
		case parser.Snowflake:
			query = "CREATE TABLE mytable (id INT, amount BIGINT);"
		case parser.Sqlite:
			query = "CREATE TABLE employees (id INT, name TEXT)"
		default:
			query = "CREATE TABLE employees (id INT, name TEXT)"
		}

		if query != "" {
			res, err := p.Parse(query, opt)
			if err != nil {
				t.Error(fmt.Sprintf("%s failed", db), err)
			}
			if res != nil && res.AST.Len() > 1 {
				t.Error("AST length is greater than 1")
			}

			for _, ast := range res.AST {
				if ast == nil {
					t.Error("AST is nil")
				} else {
					if !ast.Is(types.CreateStatement) {
						t.Error("AST type is not create")
					} else {
						createStmt, err := ast.Create()
						if err != nil {
							t.Error("Error decoding create statement:", err)
						}
						if createStmt == nil {
							t.Error("Drop statement is nil")
						}
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
		opt := parser.WithDatabase(db)
		res, err := p.Parse("DROP TABLE employees", opt)
		if err != nil {
			t.Error(fmt.Sprintf("%s failed", db), err)
		}
		if res != nil && res.AST.Len() > 1 {
			t.Error("AST length is greater than 1")
		}

		for _, ast := range res.AST {
			if ast == nil {
				t.Error("AST is nil")
			} else {
				if !ast.Is(types.DropStatement) {
					t.Error("AST type is not drop")
				} else {
					dropStmt, err := ast.Drop()
					if err != nil {
						t.Error("Error decoding drop statement:", err)
					}
					if dropStmt == nil {
						t.Error("Drop statement is nil")
					}

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

	db := parser.BigQuery
	sql := fmt.Sprintf(`
		WITH t1 AS (
			SELECT %s AS col1 FROM %s.%s
		)

		SELECT * FROM t1
	`, "test_column", "cool_schema", "test_table")
	res, err := p.Parse(sql, parser.WithDatabase(db))
	if err != nil {
		t.Error(fmt.Sprintf("%s failed", db), err)
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
	} else if res.AST.Len() == 0 {
		t.Error("AST is empty")
	} else {
		for _, ast := range res.AST {
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
