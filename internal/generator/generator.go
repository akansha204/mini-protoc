package generator

import (
	"fmt"
	"strings"

	"github.com/akansha204/mini-protoc/internal/ast"
)

type Generator struct {
	builder strings.Builder
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(proto *ast.ProtoFile) string {
	g.generatePackage(proto.Package)

	for _, msg := range proto.Messages {
		g.generateMessage(msg)
	}

	return g.builder.String()
}

func (g *Generator) generatePackage(pkg string) {
	pkg = strings.ReplaceAll(pkg, ".", "_")
	g.builder.WriteString(fmt.Sprintf("package %s\n\n", pkg))
}
