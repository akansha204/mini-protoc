package main

import (
	"fmt"

	"github.com/akansha204/mini-protoc/internal/ast"
	"github.com/akansha204/mini-protoc/internal/generator"
	"github.com/akansha204/mini-protoc/internal/validator"
)

func main() {
	proto := &ast.ProtoFile{
		Syntax:  "proto3",
		Package: "user",

		Messages: []*ast.Message{
			{
				Name: "User",
				Fields: []*ast.Field{
					{
						Type:   "string",
						Name:   "name",
						Number: 1,
					},
					{
						Type:   "int32",
						Name:   "age",
						Number: 2,
					},
				},
			},
			{
				Name: "Person",
				Fields: []*ast.Field{
					{
						Type:   "string",
						Name:   "name",
						Number: 1,
					},
					{
						Type:   "int32",
						Name:   "age",
						Number: 2,
					},
				},
			},
		},
	}

	v := validator.New()
	v.ValidateProtoFile(proto)

	if len(v.Errors()) > 0 {
		for _, err := range v.Errors() {
			fmt.Println(err)
		}
		return
	}

	gen := generator.New()
	output := gen.Generate(proto)

	fmt.Println(output)
}
