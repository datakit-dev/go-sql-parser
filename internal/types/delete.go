package types

// export interface Delete {
// 	type: "delete";
// 	table: any;
// 	from: Array<From | Dual>;
// 	where: Expr | Function | null;
// 	loc?: LocationRange;
//   }

type Delete struct {
	Type  Statement `json:"type"`
	Table any       `json:"table"`
	From  []any     `json:"from"`
	Where any       `json:"where"`
	// Loc   LocationRange `json:"loc"`
}
