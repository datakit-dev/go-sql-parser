package types

import "github.com/mitchellh/mapstructure"

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
	Database     *string       `json:"database,omitempty"`
	Type         *string       `json:"type,omitempty"`
	TrimQuery    *bool         `json:"trimQuery,omitempty"`
	ParseOptions *ParseOptions `json:"parseOptions,omitempty"`
}

type ParseOptions struct {
	IncludeLocations bool
}

func (o *Option) ToMap() (map[string]any, error) {
	m := make(map[string]any)
	err := mapstructure.Decode(o, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (o *Option) SetDatabase(db string) {
	o.Database = &db
}

func (o *Option) SetType(t string) {
	o.Type = &t
}

func (o *Option) SetTrimQuery(b bool) {
	o.TrimQuery = &b
}

func (o *Option) SetParseOptions(po ParseOptions) {
	o.ParseOptions = &po
}
