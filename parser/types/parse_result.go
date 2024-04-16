package types

type ParseResult struct {
	TableList  []string `json:"tableList"`
	ColumnList []string `json:"columnList"`
	AST        ASTs
}
