package parser

import "github.com/datakit-dev/go-sql-parser/parser/types"

type Option interface {
	Opt(types.Option)
}

type dbOpt struct {
	DB Database
}

func (opt dbOpt) Opt(o types.Option) {
	o.SetDatabase(opt.DB.String())
}

func WithDatabase(db Database) Option {
	return dbOpt{db}
}