package types

// export interface Insert {
// 	type: "insert";
// 	db: string | null;
// 	table: any;
// 	columns: string[] | null;
// 	values: InsertReplaceValue[];
// 	loc?: LocationRange;
//   }

// export interface InsertValue {
// 	type: "expr_list";
// 	value: any[];
// 	loc?: LocationRange;
//   }

type InsertValue struct {
	Type  string `json:"type"`
	Value []any  `json:"value"`
	// Loc   LocationRange `json:"loc"`
}

type Insert struct {
	Type    Statement     `json:"type"`
	DB      string        `json:"db"`
	Table   any           `json:"table"`
	Columns []string      `json:"columns"`
	Values  []InsertValue `json:"values"`
	// Loc     LocationRange `json:"loc"`
}
