package internal

import (
	"fmt"

	_ "embed"

	"github.com/datakit-dev/go-sql-parser/internal/types"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type Parser struct {
	vm *goja.Runtime
	p  *goja.Object
}

func NewParser() (*Parser, error) {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	regOpts := []require.Option{
		require.WithGlobalFolders("../js"),
	}
	reg := require.NewRegistry(regOpts...)
	req := reg.Enable(vm)

	mod, err := req.Require("index.umd.js")
	if err != nil {
		return nil, err
	}

	m := mod.ToObject(vm)
	if m == nil {
		return nil, fmt.Errorf("module not an object")
	}

	newParser, ok := goja.AssertConstructor(m.Get("Parser"))
	if !ok {
		return nil, fmt.Errorf("module not a constructor")
	}

	p, err := newParser(nil)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, fmt.Errorf("constructor returned nil")
	}

	return &Parser{vm, p}, nil
}

// parse(sql: string, opt?: Option): TableColumnAst;
func (p *Parser) Parse(sql string, opt types.Option) (*ParseResult, error) {
	if p.vm == nil {
		return nil, fmt.Errorf("vm not initialized")
	}

	v := p.p.Get("parse")
	if v == nil {
		return nil, fmt.Errorf("unknown parse failure")
	}

	parse, ok := goja.AssertFunction(v)
	if !ok {
		return nil, fmt.Errorf("unknown parse failure")
	}

	val, err := parse(p.p, p.vm.ToValue(sql), p.vm.ToValue(opt))
	if err != nil {
		return nil, err
	}

	return NewParseResult(p.vm, val)
}

// astify(sql: string, opt?: Option): AST[] | AST;
func (p *Parser) Astify(sql string, opt types.Option) (*ASTResult, error) {
	if p.vm == nil {
		return nil, fmt.Errorf("vm not initialized")
	}

	v := p.p.Get("astify")
	if v == nil {
		return nil, fmt.Errorf("unknown astify failure")
	}

	astify, ok := goja.AssertFunction(v)
	if !ok {
		return nil, fmt.Errorf("unknown astify failure")
	}

	val, err := astify(p.p, p.vm.ToValue(sql), p.vm.ToValue(opt))
	if err != nil {
		return nil, err
	}

	return NewAST(p.vm, val)
}

// sqlify(ast: AST[] | AST, opt?: Option): string;
func (p *Parser) Sqlify(ast *ASTResult, opt types.Option) (string, error) {
	if p.vm == nil {
		return "", fmt.Errorf("vm not initialized")
	}

	v := p.p.Get("sqlify")
	if v == nil {
		return "", fmt.Errorf("unknown sqlify failure")
	}

	sqlify, ok := goja.AssertFunction(v)
	if !ok {
		return "", fmt.Errorf("unknown sqlify failure")
	}

	val, err := sqlify(p.p, ast.v, p.vm.ToValue(opt))
	if err != nil {
		return "", err
	}

	return val.String(), nil
}

// exprToSQL(ast: any, opt?: Option): string;
// whiteListCheck(sql: string, whiteList: string[], opt?: Option): Error | undefined;
// tableList(sql: string, opt?: Option): string[];
// columnList(sql: string, opt?: Option): string[];
