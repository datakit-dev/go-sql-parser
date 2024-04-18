package types

// export interface Select {
// 	with: With[] | null;
// 	type: "select";
// 	options: any[] | null;
// 	distinct: "DISTINCT" | null;
// 	columns: any[] | Column[];
// 	from: From[] | null;
// 	where: Expr | Function | null;
// 	groupby: ColumnRef[] | null;
// 	having: any[] | null;
// 	orderby: OrderBy[] | null;
// 	limit: Limit | null;
// 	_orderby?: OrderBy[] | null;
// 	_limit?: Limit | null;
// 	parentheses_symbol?: boolean;
// 	_parentheses?: boolean;
// 	loc?: LocationRange;
//   }

type Select struct {
	Type    Statement `json:"type"`
	With    []With    `json:"with"`
	Options []any     `json:"options"`
	// Distinct string    `json:"distinct"`
	Columns []any `json:"columns"`
	// From               []From     `json:"from"`
	Where any `json:"where"`
	// Groupby            []ColumnRef `json:"groupby"`
	Having []any `json:"having"`
	// Orderby            []OrderBy  `json:"orderby"`
	// Limit Limit `json:"limit"`
	// Orderby_           []OrderBy  `json:"_orderby"`
	// Limit_            Limit `json:"_limit"`
	ParenthesesSymbol bool `json:"parentheses_symbol"`
	Parentheses       bool `json:"_parentheses"`
	// Loc                LocationRange `json:"loc"`
}

// export interface With {
// 	name: { value: string };
// 	stmt: {
// 	  _parentheses?: boolean;
// 	  tableList: string[];
// 	  columnList: string[];
// 	  ast: Select;
// 	};
// 	columns?: any[];
//   }

type With struct {
	Name struct {
		Value string `json:"value"`
	} `json:"name"`
	Stmt struct {
		Parentheses *bool    `json:"_parentheses"`
		TableList   []string `json:"tableList"`
		ColumnList  []string `json:"columnList"`
		Ast         Select   `json:"ast"`
	} `json:"stmt"`
	Columns []string `json:"columns"`
}
