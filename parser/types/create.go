package types

// export interface Create {
// 	type: "create";
// 	keyword: "table" | "index" | "database";
// 	temporary?: "temporary" | null;
// 	table?: { db: string; table: string }[];
// 	if_not_exists?: "if not exists" | null;
// 	like?: {
// 	  type: "like";
// 	  table: string;
// 	  parentheses?: boolean;
// 	} | null;
// 	ignore_replace?: "ignore" | "replace" | null;
// 	as?: string | null;
// 	query_expr?: any | null;
// 	create_definitions?: CreateDefinition[] | null;
// 	table_options?: any[] | null;
// 	index_using?: {
// 	  keyword: "using";
// 	  type: "btree" | "hash";
// 	} | null;
// 	index?: string | null;
// 	on_kw?: "on" | null;
// 	index_columns?: any[] | null;
// 	index_type?: "unique" | "fulltext" | "spatial" | null;
// 	index_options?: any[] | null;
// 	algorithm_option?: {
// 	  type: "alter";
// 	  keyword: "algorithm";
// 	  resource: "algorithm";
// 	  symbol: "=" | null;
// 	  algorithm: "default" | "instant" | "inplace" | "copy";
// 	} | null;
// 	lock_option?: {
// 	  type: "alter";
// 	  keyword: "lock";
// 	  resource: "lock";
// 	  symbol: "=" | null;
// 	  lock: "default" | "none" | "shared" | "exclusive";
// 	} | null;
// 	database?: string;
// 	loc?: LocationRange;
//   }

type Create struct {
	Type      Statement `json:"type"`
	Keyword   string    `json:"keyword"`
	Temporary string    `json:"temporary"`
	Table     []struct {
		DB    string `json:"db"`
		Table string `json:"table"`
	} `json:"table"`
	IfNotExists string `json:"if_not_exists"`
	Like        struct {
		Type        string `json:"type"`
		Table       string `json:"table"`
		Parentheses bool   `json:"parentheses"`
	} `json:"like"`
	IgnoreReplace     string `json:"ignore_replace"`
	As                string `json:"as"`
	QueryExpr         any    `json:"query_expr"`
	CreateDefinitions []any  `json:"create_definitions"`
	TableOptions      []any  `json:"table_options"`
	IndexUsing        struct {
		Keyword string `json:"keyword"`
		Type    string `json:"type"`
	} `json:"index_using"`
	Index           string `json:"index"`
	OnKw            string `json:"on_kw"`
	IndexColumns    []any  `json:"index_columns"`
	IndexType       string `json:"index_type"`
	IndexOptions    []any  `json:"index_options"`
	AlgorithmOption struct {
		Type      string `json:"type"`
		Keyword   string `json:"keyword"`
		Resource  string `json:"resource"`
		Symbol    string `json:"symbol"`
		Algorithm string `json:"algorithm"`
	} `json:"algorithm_option"`
	LockOption struct {
		Type     string `json:"type"`
		Keyword  string `json:"keyword"`
		Resource string `json:"resource"`
		Symbol   string `json:"symbol"`
		Lock     string `json:"lock"`
	} `json:"lock_option"`
	Database string `json:"database"`
	// Loc                LocationRange `json:"loc"`
}
