package internal

import "github.com/dop251/goja_nodejs/require"

var (
	registry *require.Registry
)

func init() {
	regOpts := []require.Option{
		require.WithGlobalFolders("../js"),
	}
	registry = require.NewRegistry(regOpts...)
}
