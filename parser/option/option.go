package option

import (
	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/internal/types"
)

type Option interface {
	Opt(types.Option)
}

type dbOpt struct {
	DB internal.Database
}

func (opt dbOpt) Opt(o types.Option) {
	o.SetDatabase(opt.DB.String())
}

func WithDatabase(db internal.Database) Option {
	return dbOpt{db}
}
