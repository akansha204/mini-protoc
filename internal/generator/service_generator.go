package generator

import (
	"fmt"

	"github.com/akansha204/mini-protoc/internal/ast"
)

func (g *Generator) generateService(service *ast.Service) {
	fmt.Fprintf(
		&g.builder,
		"type %s interface {\n",
		service.Name,
	)
	for _, rpc := range service.RPC {
		fmt.Fprintf(
			&g.builder,
			"\t%s(req %s) (%s, error)\n",
			rpc.Name,
			rpc.RequestType,
			rpc.ResponseType,
		)
	}

	g.builder.WriteString("}\n\n")
}
