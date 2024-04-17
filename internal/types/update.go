package types

// export interface Update {
// 	type: "update";
// 	db: string | null;
// 	table: Array<From | Dual> | null;
// 	set: SetList[];
// 	where: Expr | Function | null;
// 	loc?: LocationRange;
//   }

type Update struct {
	Type  Statement `json:"type"`
	DB    string    `json:"db"`
	Table []any     `json:"table"`
	// Set   []SetList `json:"set"`
	Where any `json:"where"`
	// Loc   LocationRange `json:"loc"`
}
