package types

// export interface Use {
// 	type: "use";
// 	db: string;
// 	loc?: LocationRange;
//   }

type Use struct {
	Type Statement `json:"type"`
	DB   string    `json:"db"`
	// Loc  LocationRange `json:"loc"`
}
