package parser

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/internal"
	"github.com/datakit-dev/go-sql-parser/parser/types"
)

var (
	ErrParserNotInitialized = fmt.Errorf("parser not initialized")
)

// type TableColumnList *internal.TableColumnList

// type ParseResult *internal.ParseResult

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

func (p *Parser) Parse(sql string, opts ...Option) (*ParseResult, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	res, err := p.p.Parse(sql, opt)
	if err != nil {
		return nil, err
	}
	return NewParseResult(res), nil
}

func (p *Parser) Astify(sql string, opts ...Option) (*ASTResult, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	ast, err := p.p.Astify(sql, opt)
	if err != nil {
		return nil, err
	}
	return NewASTResult(ast), nil
}

func (p *Parser) Sqlify(ast *ASTResult, opts ...Option) (string, error) {
	if err := p.preCheck(); err != nil {
		return "", err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.Sqlify(ast.ASTResult, opt)
}

func (p *Parser) TableList(sql string, opts ...Option) ([]string, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.TableList(sql, opt)
}

func (p *Parser) ColumnList(sql string, opts ...Option) ([]string, error) {
	if err := p.preCheck(); err != nil {
		return nil, err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.ColumnList(sql, opt)
}

func (p *Parser) WhiteListCheck(sql string, whiteList []string, opts ...Option) error {
	if err := p.preCheck(); err != nil {
		return err
	}
	opt := &types.Option{}
	for _, o := range opts {
		o.Opt(opt)
	}
	return p.p.WhiteListCheck(sql, whiteList, opt)
}
