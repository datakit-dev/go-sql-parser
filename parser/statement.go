package parser

import "github.com/datakit-dev/go-sql-parser/internal/types"

const (
	UnknownStatement = types.UnknownStatement
	UseStatement     = types.UseStatement
	SelectStatement  = types.SelectStatement
	InsertStatement  = types.InsertStatement
	ReplaceStatement = types.ReplaceStatement
	UpdateStatement  = types.UpdateStatement
	DeleteStatement  = types.DeleteStatement
	AlterStatement   = types.AlterStatement
	CreateStatement  = types.CreateStatement
	DropStatement    = types.DropStatement
)
