package parser

import "github.com/dop251/goja"

type Option interface {
	Opt(*goja.Object)
}

type dbOpt struct {
	DB Database
}

func (o dbOpt) Opt(obj *goja.Object) {
	obj.Set("database", o.DB.String())
}

func WithDatabase(db Database) Option {
	return dbOpt{DB: db}
}
