package parser

import "github.com/datakit-dev/go-sql-parser/internal"

type ParseResult struct {
	AST ASTs
	internal.TableColumnList
}

func NewParseResult(val *internal.ParseResult) *ParseResult {
	return &ParseResult{NewAST(val.ASTResult), val.TableColumnList}
}
