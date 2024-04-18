package types

// export interface BaseFrom {
// 	db: string | null;
// 	table: string;
// 	as: string | null;
// 	schema?: string;
// 	loc?: LocationRange;
// }

// export interface Join extends BaseFrom {
// 	join: "INNER JOIN" | "LEFT JOIN" | "RIGHT JOIN";
// 	using?: string[];
// 	on?: Expr;
// }

// export interface TableExpr {
// 	expr: {
// 		ast: Select;
// 	};
// 	as?: string | null;
// }

// export interface Dual {
// 	type: "dual";
// 	loc?: LocationRange;
// }

// export type From = BaseFrom | Join | TableExpr | Dual;

type BaseFrom struct {
	DB     *string `json:"db,omitempty"`
	Schema *string `json:"schema,omitempty"`
	As     *string `json:"as,omitempty"`
	Table  string  `json:"table"`
	// Loc    LocationRange `json:"loc"`
}

type Join struct {
	BaseFrom
	Join  string   `json:"join"`
	Using []string `json:"using,omitempty"`
	On    any      `json:"on,omitempty"`
}

type TableExpr struct {
	Expr struct {
		Ast Select `json:"ast"`
	} `json:"expr"`
	As *string `json:"as,omitempty"`
}

type Dual struct {
	Type string `json:"type"`
	// Loc  LocationRange `json:"loc"`
}

type From struct {
	BaseFrom
	Join      *Join      `json:"join,omitempty"`
	TableExpr *TableExpr `json:"tableExpr,omitempty"`
	Dual      *Dual      `json:"dual,omitempty"`
}
