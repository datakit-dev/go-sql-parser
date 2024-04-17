package types

// export interface Drop {
// 	type: "drop";
// 	keyword: string;
// 	name: any[];
//   }

type Drop struct {
	Type    Statement `json:"type"`
	Keyword string    `json:"keyword"`
	Name    []any     `json:"name"`
}
