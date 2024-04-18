package parser

import "github.com/datakit-dev/go-sql-parser/internal"

type ParseResult struct {
	*ASTResult
	internal.TableColumnList
}

func NewParseResult(val *internal.ParseResult) *ParseResult {
	return &ParseResult{NewASTResult(val.ASTResult), val.TableColumnList}
}
