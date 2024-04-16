package parser

import (
	"fmt"

	"github.com/datakit-dev/go-sql-parser/parser/types"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type Parser struct {
	vm *goja.Runtime
	p  *goja.Object
}

func New() (*Parser, error) {
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

func (p *Parser) parse(sql string, opts ...Option) (val goja.Value, err error) {
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

	opt := p.vm.NewObject()
	if len(opts) > 0 {
		for _, o := range opts {
			o.Opt(opt)
		}
		val, err = parse(p.p, p.vm.ToValue(sql), p.vm.ToValue(opt))
	} else {
		val, err = parse(p.p, p.vm.ToValue(sql))
	}
	if err != nil {
		return nil, err
	}

	return
}

func (p *Parser) Parse(sql string, opts ...Option) (*types.ParseResult, error) {
	res := types.ParseResult{}
	val, err := p.parse(sql, opts...)
	if err != nil {
		return nil, err
	}

	err = p.vm.ExportTo(val, &res)
	if err != nil {
		return nil, err
	}

	rObj := val.ToObject(p.vm)
	if rObj == nil {
		return nil, fmt.Errorf("parse failed")
	}

	if ast := rObj.Get("ast"); ast != nil {
		if _, ok := ast.Export().([]any); ok {
			err = p.vm.ExportTo(ast, &res.AST)
			if err != nil {
				return nil, err
			}
		} else if _, ok := ast.Export().(map[string]any); ok {
			var a map[string]any
			err = p.vm.ExportTo(ast, &a)
			if err != nil {
				return nil, err
			}
			res.AST = append(res.AST, a)
		}
	} else {
		return nil, fmt.Errorf("AST not found")
	}

	return &res, nil
}
