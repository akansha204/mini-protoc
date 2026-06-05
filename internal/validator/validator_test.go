package validator

import (
	"strings"
	"testing"

	"github.com/akansha204/mini-protoc/internal/ast"
	"github.com/akansha204/mini-protoc/internal/lexer"
	"github.com/akansha204/mini-protoc/internal/parser"
)

func TestValidateAcceptsValidProtoFile(t *testing.T) {
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

		message Envelope {
			AddRequest request = 1;
			string trace_id = 2;
			bool active = 3;
			int64 count = 4;
			float score = 5;
			double ratio = 6;
		}

		service CalculatorService {
			rpc Add(AddRequest) returns (AddResponse);
		}
	`

	file := parseProtoFile(t, input)

	v := New()
	v.ValidateProtoFile(file)

	if len(v.Errors()) != 0 {
		t.Fatalf("expected no validation errors, got %v", v.Errors())
	}
}

func TestValidateRejectsDuplicateMessageNames(t *testing.T) {
	file := validFile()
	file.Messages = append(file.Messages, &ast.Message{Name: "AddRequest"})

	assertValidationError(t, file, "duplicate message name: AddRequest")
}

func TestValidateRejectsInvalidFields(t *testing.T) {
	tests := []struct {
		name      string
		fields    []*ast.Field
		wantError string
	}{
		{
			name: "duplicate field names",
			fields: []*ast.Field{
				{Type: "int32", Name: "id", Number: 1},
				{Type: "string", Name: "id", Number: 2},
			},
			wantError: "duplicate field name 'id' in message AddRequest",
		},
		{
			name: "duplicate field numbers",
			fields: []*ast.Field{
				{Type: "int32", Name: "id", Number: 1},
				{Type: "string", Name: "name", Number: 1},
			},
			wantError: "duplicate field number 1 in message AddRequest",
		},
		{
			name: "missing field name",
			fields: []*ast.Field{
				{Type: "int32", Number: 1},
			},
			wantError: "field name is required in message AddRequest",
		},
		{
			name: "missing field type",
			fields: []*ast.Field{
				{Name: "id", Number: 1},
			},
			wantError: "field type is required in message AddRequest",
		},
		{
			name: "unsupported primitive type",
			fields: []*ast.Field{
				{Type: "bytes", Name: "payload", Number: 1},
			},
			wantError: "unknown type 'bytes' in message AddRequest",
		},
		{
			name: "unknown message reference",
			fields: []*ast.Field{
				{Type: "MissingMessage", Name: "missing", Number: 1},
			},
			wantError: "unknown type 'MissingMessage' in message AddRequest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := validFile()
			file.Messages[0].Fields = tt.fields

			assertValidationError(t, file, tt.wantError)
		})
	}
}

func TestValidateRejectsDuplicateServiceNames(t *testing.T) {
	file := validFile()
	file.Services = append(
		file.Services,
		&ast.Service{Name: "CalculatorService"},
	)

	assertValidationError(t, file, "duplicate service name: CalculatorService")
}

func TestValidateRejectsInvalidRPCs(t *testing.T) {
	tests := []struct {
		name      string
		rpcs      []*ast.RPC
		wantError string
	}{
		{
			name: "duplicate rpc names",
			rpcs: []*ast.RPC{
				{Name: "Add", RequestType: "AddRequest", ResponseType: "AddResponse"},
				{Name: "Add", RequestType: "AddRequest", ResponseType: "AddResponse"},
			},
			wantError: "duplicate rpc name 'Add' in service CalculatorService",
		},
		{
			name: "missing request type",
			rpcs: []*ast.RPC{
				{Name: "Add", ResponseType: "AddResponse"},
			},
			wantError: "request type is required in rpc Add",
		},
		{
			name: "missing response type",
			rpcs: []*ast.RPC{
				{Name: "Add", RequestType: "AddRequest"},
			},
			wantError: "response type is required in rpc Add",
		},
		{
			name: "unknown request type",
			rpcs: []*ast.RPC{
				{Name: "Add", RequestType: "MissingRequest", ResponseType: "AddResponse"},
			},
			wantError: "unknown request type 'MissingRequest' in rpc Add",
		},
		{
			name: "unknown response type",
			rpcs: []*ast.RPC{
				{Name: "Add", RequestType: "AddRequest", ResponseType: "MissingResponse"},
			},
			wantError: "unknown response type 'MissingResponse' in rpc Add",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := validFile()
			file.Services[0].RPC = tt.rpcs

			assertValidationError(t, file, tt.wantError)
		})
	}
}

func TestValidateClearsErrorsBetweenRuns(t *testing.T) {
	v := New()
	v.ValidateProtoFile(&ast.ProtoFile{})

	if len(v.Errors()) == 0 {
		t.Fatal("expected invalid proto file to produce errors")
	}

	v.ValidateProtoFile(validFile())

	if len(v.Errors()) != 0 {
		t.Fatalf("expected errors to be cleared, got %v", v.Errors())
	}
}

func parseProtoFile(t *testing.T, input string) *ast.ProtoFile {
	t.Helper()

	l := lexer.New(input)
	p := parser.New(l)
	file := p.ParseProtoFile()

	if len(p.Errors()) != 0 {
		t.Fatalf("parser had errors: %v", p.Errors())
	}

	return file
}

func validFile() *ast.ProtoFile {
	return &ast.ProtoFile{
		Syntax:  "proto3",
		Package: "calculator.math",
		Messages: []*ast.Message{
			{
				Name: "AddRequest",
				Fields: []*ast.Field{
					{Type: "int32", Name: "a", Number: 1},
					{Type: "int32", Name: "b", Number: 2},
				},
			},
			{
				Name: "AddResponse",
				Fields: []*ast.Field{
					{Type: "int32", Name: "result", Number: 1},
				},
			},
		},
		Services: []*ast.Service{
			{
				Name: "CalculatorService",
				RPC: []*ast.RPC{
					{Name: "Add", RequestType: "AddRequest", ResponseType: "AddResponse"},
				},
			},
		},
	}
}

func assertValidationError(
	t *testing.T,
	file *ast.ProtoFile,
	wantError string,
) {
	t.Helper()

	v := New()
	v.ValidateProtoFile(file)

	for _, err := range v.Errors() {
		if strings.Contains(err, wantError) {
			return
		}
	}

	t.Fatalf("expected error %q, got %v", wantError, v.Errors())
}
