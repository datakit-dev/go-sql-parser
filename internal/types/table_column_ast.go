package types

// export interface TableColumnAst {
// 	tableList: string[];
// 	columnList: string[];
// 	ast: AST[] | AST;
// 	loc?: LocationRange;
//   }

type TableColumnAst struct {
	TableList  []string `json:"tableList"`
	ColumnList []string `json:"columnList"`
	// AST        any      `json:"ast"`
	// Loc *LocationRange
}
