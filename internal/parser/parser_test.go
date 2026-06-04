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
		 message AddRequest {
	        int32 a = 1;
	        int64 b = 2;
	        string user_name = 3;
	        bool active = 4;
	        float x = 5;
	        double y = 6;
        }
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if len(file.Messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(file.Messages))
	}

	msg := file.Messages[0]

	if msg.Name != "AddRequest" {
		t.Fatalf("expected message name AddRequest, got %q", msg.Name)
	}
	if len(msg.Fields) != 6 {
		t.Fatalf("expected 6 fields, got %d", len(msg.Fields))
	}

	tests := []struct {
		fieldType string
		fieldName string
		number    int
	}{
		{"int32", "a", 1},
		{"int64", "b", 2},
		{"string", "user_name", 3},
		{"bool", "active", 4},
		{"float", "x", 5},
		{"double", "y", 6},
	}

	for i, tt := range tests {
		field := msg.Fields[i]

		if field.Type != tt.fieldType {
			t.Fatalf(
				"field[%d] type wrong. expected=%q got=%q",
				i,
				tt.fieldType,
				field.Type,
			)
		}

		if field.Name != tt.fieldName {
			t.Fatalf(
				"field[%d] name wrong. expected=%q got=%q",
				i,
				tt.fieldName,
				field.Name,
			)
		}

		if field.Number != tt.number {
			t.Fatalf(
				"field[%d] number wrong. expected=%d got=%d",
				i,
				tt.number,
				field.Number,
			)
		}
	}

}
func TestParseService(t *testing.T) {
	input := `
		service CalculatorService {
	        rpc Add(AddRequest) returns (AddResponse);
	        rpc Multiply(AddRequest) returns (AddResponse);
        } 
	`
	l := lexer.New(input)
	p := New(l)
	file := p.ParseProtoFile()
	if len(file.Services) != 1 {
		t.Fatalf("expected 1 service, got %d", len(file.Services))
	}

	svc := file.Services[0]

	if svc.Name != "CalculatorService" {
		t.Fatalf(
			"expected CalculatorService, got %q",
			svc.Name,
		)
	}

	if len(svc.RPC) != 2 {
		t.Fatalf(
			"expected 2 RPCs, got %d",
			len(svc.RPC),
		)
	}

	tests := []struct {
		name     string
		request  string
		response string
	}{
		{"Add", "AddRequest", "AddResponse"},
		{"Multiply", "AddRequest", "AddResponse"},
	}

	for i, tt := range tests {
		rpc := svc.RPC[i]

		if rpc.Name != tt.name {
			t.Fatalf(
				"rpc[%d] name wrong. expected=%q got=%q",
				i,
				tt.name,
				rpc.Name,
			)
		}

		if rpc.RequestType != tt.request {
			t.Fatalf(
				"rpc[%d] request wrong. expected=%q got=%q",
				i,
				tt.request,
				rpc.RequestType,
			)
		}

		if rpc.ResponseType != tt.response {
			t.Fatalf(
				"rpc[%d] response wrong. expected=%q got=%q",
				i,
				tt.response,
				rpc.ResponseType,
			)
		}
	}
}

func TestParseCompleteProtoFile(t *testing.T) {
	input := `
		syntax = "proto3";

		package calculator.math;

		message AddRequest {
			int32 a = 1;
			int32 b = 2;
		}

		message AddResponse {
			int32 result = 1;
		}

		service CalculatorService {
			rpc Add(AddRequest) returns (AddResponse);
		}
	`

	l := lexer.New(input)
	p := New(l)

	file := p.ParseProtoFile()

	if len(p.Errors()) != 0 {
		t.Fatalf("parser had errors: %v", p.Errors())
	}

	// syntax
	if file.Syntax != "proto3" {
		t.Fatalf(
			"wrong syntax. expected=%q got=%q",
			"proto3",
			file.Syntax,
		)
	}

	// package
	if file.Package != "calculator.math" {
		t.Fatalf(
			"wrong package. expected=%q got=%q",
			"calculator.math",
			file.Package,
		)
	}

	// messages
	if len(file.Messages) != 2 {
		t.Fatalf(
			"expected 2 messages, got %d",
			len(file.Messages),
		)
	}

	// AddRequest
	req := file.Messages[0]

	if req.Name != "AddRequest" {
		t.Fatalf(
			"wrong message name. expected=%q got=%q",
			"AddRequest",
			req.Name,
		)
	}

	if len(req.Fields) != 2 {
		t.Fatalf(
			"expected 2 fields, got %d",
			len(req.Fields),
		)
	}

	if req.Fields[0].Type != "int32" ||
		req.Fields[0].Name != "a" ||
		req.Fields[0].Number != 1 {
		t.Fatalf("first field parsed incorrectly")
	}

	if req.Fields[1].Type != "int32" ||
		req.Fields[1].Name != "b" ||
		req.Fields[1].Number != 2 {
		t.Fatalf("second field parsed incorrectly")
	}

	// AddResponse
	resp := file.Messages[1]

	if resp.Name != "AddResponse" {
		t.Fatalf(
			"wrong message name. expected=%q got=%q",
			"AddResponse",
			resp.Name,
		)
	}

	if len(resp.Fields) != 1 {
		t.Fatalf(
			"expected 1 field, got %d",
			len(resp.Fields),
		)
	}

	if resp.Fields[0].Name != "result" {
		t.Fatalf(
			"wrong field name. expected=%q got=%q",
			"result",
			resp.Fields[0].Name,
		)
	}

	// service
	if len(file.Services) != 1 {
		t.Fatalf(
			"expected 1 service, got %d",
			len(file.Services),
		)
	}

	svc := file.Services[0]

	if svc.Name != "CalculatorService" {
		t.Fatalf(
			"wrong service name. expected=%q got=%q",
			"CalculatorService",
			svc.Name,
		)
	}

	if len(svc.RPC) != 1 {
		t.Fatalf(
			"expected 1 rpc, got %d",
			len(svc.RPC),
		)
	}

	rpc := svc.RPC[0]

	if rpc.Name != "Add" {
		t.Fatalf(
			"wrong rpc name. expected=%q got=%q",
			"Add",
			rpc.Name,
		)
	}

	if rpc.RequestType != "AddRequest" {
		t.Fatalf(
			"wrong request type. expected=%q got=%q",
			"AddRequest",
			rpc.RequestType,
		)
	}

	if rpc.ResponseType != "AddResponse" {
		t.Fatalf(
			"wrong response type. expected=%q got=%q",
			"AddResponse",
			rpc.ResponseType,
		)
	}
}
