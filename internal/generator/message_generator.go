package generator

import (
	"fmt"
	"strings"

	"github.com/akansha204/mini-protoc/internal/ast"
)

func (g *Generator) generateMessage(msg *ast.Message) {
	fmt.Fprintf(
		&g.builder,
		"type %s struct {\n",
		msg.Name,
	)

	for _, field := range msg.Fields {
		fmt.Fprintf(
			&g.builder,
			"\t%s %s\n",
			exportName(field.Name),
			goType(field.Type),
		)
	}
	g.builder.WriteString("}\n\n")
}

func exportName(name string) string {
	if len(name) == 0 {
		return ""
	}

	return strings.ToUpper(name[:1]) + name[1:]
}
