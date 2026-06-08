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

	fmt.Fprintf(
		&g.builder,
		"type %sClient struct {\n",
		service.Name,
	)

	g.builder.WriteString("}\n\n")

	fmt.Fprintf(
		&g.builder,
		"func New%sClient() *%sClient {\n",
		service.Name,
		service.Name,
	)

	fmt.Fprintf(
		&g.builder,
		"\treturn &%sClient{}\n",
		service.Name,
	)

	g.builder.WriteString("}\n\n")

	for _, rpc := range service.RPC {
		fmt.Fprintf(
			&g.builder,
			"func (c *%sClient) %s(req %s) (%s, error) {\n",
			service.Name,
			rpc.Name,
			rpc.RequestType,
			rpc.ResponseType,
		)

		g.builder.WriteString("\tpanic(\"not implemented\")\n")

		g.builder.WriteString("}\n\n")
	}

	fmt.Fprintf(
		&g.builder,
		"func Register%s(service %s) {\n",
		service.Name,
		service.Name,
	)

	g.builder.WriteString("\tpanic(\"not implemented\")\n")

	g.builder.WriteString("}\n\n")
}
