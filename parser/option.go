package parser

import "github.com/dop251/goja"

type Option interface {
	Opt(*goja.Object)
}

type dbOpt struct {
	DB Database
}

// export interface ParseOptions {
// includeLocations?: boolean;
// }
// export interface Option {
// database?: string;
// type?: string;
// trimQuery?: boolean;
// parseOptions?: ParseOptions;
// }

func (o dbOpt) Opt(obj *goja.Object) {
	obj.Set("database", o.DB.String())
}

func WithDatabase(db Database) Option {
	return dbOpt{DB: db}
}
