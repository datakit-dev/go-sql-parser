package internal

import (
	"embed"

	"github.com/dop251/goja_nodejs/require"
)

//go:embed js/*
var js embed.FS

var (
	registry *require.Registry
)

func sourceLoader(fs embed.FS) require.SourceLoader {
	return func(path string) ([]byte, error) {
		return fs.ReadFile(path)
	}
}

func init() {
	regOpts := []require.Option{
		require.WithGlobalFolders("./js"),
		require.WithLoader(sourceLoader(js)),
	}
	registry = require.NewRegistry(regOpts...)
}
