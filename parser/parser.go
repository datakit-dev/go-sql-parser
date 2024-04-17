package parser

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/internal/types"
	"github.com/datakit-dev/go-sql-parser/parser/option"
)

var (
	ErrParserNotInitialized = fmt.Errorf("parser not initialized")
)

type Parser struct {
	p *internal.Parser
}

func New() (*Parser, error) {
	p, err := internal.NewParser()
	if err != nil {
		return nil, err
	}
	return &Parser{p}, nil
}

func (p *Parser) preCheck() error {
	if p.p == nil {
		return ErrParserNotInitialized
	}
	return nil
}

func (p *Parser) Parse(sql string, opts ...option.Option) (*internal.ParseResult, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.Parse(sql, opt)
}

func (p *Parser) Astify(sql string, opts ...option.Option) (*internal.ASTResult, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.Astify(sql, opt)
}

func (p *Parser) Sqlify(ast *internal.ASTResult, opts ...option.Option) (string, error) {
	if err := p.preCheck(); err != nil {
		return "", err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.Sqlify(ast, opt)
}

func (p *Parser) TableList(sql string, opts ...option.Option) ([]string, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.TableList(sql, opt)
}

func (p *Parser) ColumnList(sql string, opts ...option.Option) ([]string, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.ColumnList(sql, opt)
}

func (p *Parser) WhiteListCheck(sql string, whiteList []string, opts ...option.Option) error {
	if err := p.preCheck(); err != nil {
		return err
	}
	opt := types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.WhiteListCheck(sql, whiteList, opt)
}
