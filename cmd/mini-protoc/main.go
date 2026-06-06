package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akansha204/mini-protoc/internal/generator"
	"github.com/akansha204/mini-protoc/internal/lexer"
	"github.com/akansha204/mini-protoc/internal/parser"
	"github.com/akansha204/mini-protoc/internal/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mini-protoc <file.proto>")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("error reading file:", err)
		os.Exit(1)
	}

	l := lexer.New(string(data))
	p := parser.New(l)

	proto := p.ParseProtoFile()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			fmt.Println("parser error:", err)
		}
		os.Exit(1)
	}

	v := validator.New()
	v.ValidateProtoFile(proto)

	if len(v.Errors()) > 0 {
		for _, err := range v.Errors() {
			fmt.Println("validation error:", err)
		}
		os.Exit(1)
	}

	gen := generator.New()
	output := gen.Generate(proto)

	outputPath := strings.TrimSuffix(
		inputPath,
		".proto",
	) + ".pb.go"

	err = os.WriteFile(
		outputPath,
		[]byte(output),
		0644,
	)

	if err != nil {
		fmt.Println("error writing output file:", err)
		os.Exit(1)
	}

	fmt.Printf("generated %s\n", outputPath)
}
