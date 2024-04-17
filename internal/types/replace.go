package types

// export interface Replace {
// 	type: "replace";
// 	db: string | null;
// 	table: any;
// 	columns: string[] | null;
// 	values: InsertReplaceValue[];
// 	loc?: LocationRange;
//   }

// export interface ReplaceValue {
// 	type: "expr_list";
// 	value: any[];
// 	loc?: LocationRange;
//   }

type ReplaceValue struct {
	Type  string `json:"type"`
	Value []any  `json:"value"`
	// Loc   LocationRange `json:"loc"`
}

type Replace struct {
	Type    Statement      `json:"type"`
	DB      string         `json:"db"`
	Table   any            `json:"table"`
	Columns []string       `json:"columns"`
	Values  []ReplaceValue `json:"values"`
	// Loc     LocationRange `json:"loc"`
}
