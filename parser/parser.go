package parser

import (
	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/internal/types"
	"github.com/datakit-dev/go-sql-parser/parser/option"
)

type Parser struct {
	*internal.Parser
}

func New() (*Parser, error) {
	p, err := internal.NewParser()
	if err != nil {
		return nil, err
	}
	return &Parser{p}, nil
}

// parse(sql: string, opt?: Option): TableColumnAst;
func (p *Parser) Parse(sql string, opts ...option.Option) (*internal.ParseResult, error) {
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.Parser.Parse(sql, opt)
}

// astify(sql: string, opt?: Option): AST[] | AST;
func (p *Parser) Astify(sql string, opts ...option.Option) (*internal.ASTResult, error) {
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.Parser.Astify(sql, opt)
}

// sqlify(ast: AST[] | AST, opt?: Option): string;
func (p *Parser) Sqlify(ast *internal.ASTResult, opts ...option.Option) (string, error) {
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.Parser.Sqlify(ast, opt)
}

// exprToSQL(ast: any, opt?: Option): string;
// whiteListCheck(sql: string, whiteList: string[], opt?: Option): Error | undefined;
// tableList(sql: string, opt?: Option): string[];
// columnList(sql: string, opt?: Option): string[];
