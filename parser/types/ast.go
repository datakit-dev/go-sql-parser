package types

type AST map[string]any
type ASTs []AST

type Statement string

const (
	UnknownStatement Statement = ""
	UseStatement     Statement = "use"
	SelectStatement  Statement = "select"
	InsertStatement  Statement = "insert"
	ReplaceStatement Statement = "replace"
	UpdateStatement  Statement = "update"
	DeleteStatement  Statement = "delete"
	AlterStatement   Statement = "alter"
	CreateStatement  Statement = "create"
	DropStatement    Statement = "drop"
)

var Statements = []Statement{
	UseStatement,
	SelectStatement,
	InsertStatement,
	ReplaceStatement,
	UpdateStatement,
	DeleteStatement,
	AlterStatement,
	CreateStatement,
	DropStatement,
}

func (s Statement) String() string {
	return string(s)
}

func statementFrom(t string) Statement {
	for _, s := range Statements {
		if s.String() == t {
			return s
		}
	}
	return UnknownStatement
}

func (a AST) Type() Statement {
	return statementFrom(a["type"].(string))
}
