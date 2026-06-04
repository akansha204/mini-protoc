package parser

import (
	"testing"

	"github.com/akansha204/mini-protoc/internal/lexer"
)

func TestParseSyntax(t *testing.T) {
	input := `
		syntax = "proto3";
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if file.Syntax != "proto3" {
		t.Fatalf("expected syntax to be 'proto3', got %q", file.Syntax)
	}
}
func TestParsePackage(t *testing.T) {
	input := `
		package example.test.api.v1;
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if file.Package != "example.test.api.v1" {
		t.Fatalf("expected package to be 'example.test.api.v1', got %q", file.Package)
	}
}
func TestParseMessage(t *testing.T) {
	input := `
		message Person {
		}
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if file.Messages == nil {
		t.Fatalf("expected messages to be parsed")
	}
}
func TestParseService(t *testing.T) {
	input := `
		service Greeter {
		}
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if file.Services == nil {
		t.Fatalf("expected services to be parsed")
	}
}
