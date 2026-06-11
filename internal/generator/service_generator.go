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
		"func Register%s(server *rpc.Server, service %s) {\n",
		service.Name,
		service.Name,
	)

	for _, rpc := range service.RPC {

		fmt.Fprintf(
			&g.builder,
			"\tserver.Register(\n",
		)

		fmt.Fprintf(
			&g.builder,
			"\t\t%q,\n",
			service.Name+"/"+rpc.Name,
		)

		g.builder.WriteString("\t\tfunc(payload []byte) ([]byte, error) {\n\n")

		fmt.Fprintf(
			&g.builder,
			"\t\t\tvar req %s\n\n",
			rpc.RequestType,
		)

		g.builder.WriteString("\t\t\tif err := server.Decode(payload, &req); err != nil {\n")
		g.builder.WriteString("\t\t\t\treturn nil, err\n")
		g.builder.WriteString("\t\t\t}\n\n")

		fmt.Fprintf(
			&g.builder,
			"\t\t\tresp, err := service.%s(req)\n",
			rpc.Name,
		)

		g.builder.WriteString("\t\t\tif err != nil {\n")
		g.builder.WriteString("\t\t\t\treturn nil, err\n")
		g.builder.WriteString("\t\t\t}\n\n")

		g.builder.WriteString("\t\t\treturn server.Encode(resp)\n")

		g.builder.WriteString("\t\t},\n")
		g.builder.WriteString("\t)\n\n")
	}

	g.builder.WriteString("}\n\n")
}
