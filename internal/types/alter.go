package types

// export interface Alter {
// 	type: "alter";
// 	table: From[];
// 	expr: any;
// 	loc?: LocationRange;
//   }

type Alter struct {
	Type Statement `json:"type"`
	// Table []From    `json:"table"`
	Expr any `json:"expr"`
	// Loc   LocationRange `json:"loc"`
}
