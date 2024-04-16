package types

type AST map[string]any
type ASTs []AST

type Statement string

const (
	StatementUse     Statement = "use"
	StatementSelect  Statement = "select"
	StatementInsert  Statement = "insert"
	StatementReplace Statement = "replace"
	StatementUpdate  Statement = "update"
	StatementDelete  Statement = "delete"
	StatementAlter   Statement = "alter"
	StatementCreate  Statement = "create"
	StatementDrop    Statement = "drop"
)

func typeFrom(t string) Statement {
	switch t {
	case "use":
		return StatementUse
	case "select":
		return StatementSelect
	case "insert":
		return StatementInsert
	case "replace":
		return StatementReplace
	case "update":
		return StatementUpdate
	case "delete":
		return StatementDelete
	case "alter":
		return StatementAlter
	case "create":
		return StatementCreate
	case "drop":
		return StatementDrop
	}
	return ""
}

func (a AST) Type() Statement {
	return typeFrom(a["type"].(string))
}
