package types

import (
	"encoding/json"
)

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
	IncludeLocations bool `json:"includeLocations,omitempty"`
}

func (o *Option) ToMap() (map[string]any, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	m := make(map[string]any)
	if err := json.Unmarshal(b, &m); err != nil {
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

func (o *Option) SetParseOptions(po *ParseOptions) {
	o.ParseOptions = po
}
