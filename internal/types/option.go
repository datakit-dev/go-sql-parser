package types

// export interface ParseOptions {
// 	includeLocations?: boolean;
// }

// export interface Option {
// 	database?: string;
// 	type?: string;
// 	trimQuery?: boolean;
// 	parseOptions?: ParseOptions;
//   }

type Option struct {
	Database     string
	Type         string
	TrimQuery    bool
	ParseOptions ParseOptions
}

type ParseOptions struct {
	IncludeLocations bool
}

func (o *Option) SetDatabase(db string) {
	o.Database = db
}

func (o *Option) SetType(t string) {
	o.Type = t
}

func (o *Option) SetTrimQuery(b bool) {
	o.TrimQuery = b
}

func (o *Option) SetParseOptions(po ParseOptions) {
	o.ParseOptions = po
}
