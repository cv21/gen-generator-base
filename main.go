// File generated by gen. DO NOT EDIT.
// Generator plugin github.com/cv21/gen-generator-base 1.0.1
package main

import (
	"github.com/cv21/gen-generator-base/generator"

	"github.com/cv21/gen/pkg"
	plugin "github.com/hashicorp/go-plugin"
)

const (
	pluginRepoURL = "github.com/cv21/gen-generator-base@v1.0.1"
)

func main() {
	pkg.RegisterGobTypes()
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pkg.DefaultHandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginRepoURL: &pkg.NetRPCWorker{Impl: generator.NewGenerator()},
		},
	})
}
