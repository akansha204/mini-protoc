package validator

import (
	"fmt"

	"github.com/akansha204/mini-protoc/internal/ast"
)

type Validator struct {
	errors []string
}

func New() *Validator {
	return &Validator{
		errors: []string{},
	}
}

func (v *Validator) Errors() []string {
	return v.errors
}

var primitiveTypes = map[string]struct{}{
	"string": {},
	"int32":  {},
	"int64":  {},
	"bool":   {},
	"float":  {},
	"double": {},
}

func (v *Validator) ValidateProtoFile(file *ast.ProtoFile) {
	v.errors = []string{}

	if file == nil {
		v.errors = append(v.errors, "proto file is required")
		return
	}

	if file.Syntax != "proto3" {
		v.errors = append(
			v.errors,
			"invalid syntax: expected proto3",
		)
	}

	if file.Package == "" {
		v.errors = append(
			v.errors,
			"package name is required",
		)
	}

	messageTypes := make(map[string]struct{})

	for _, msg := range file.Messages {
		if msg == nil || msg.Name == "" {
			continue
		}
		messageTypes[msg.Name] = struct{}{}
	}

	v.validateMessages(file, messageTypes)
	v.validateServices(file, messageTypes)
}

func (v *Validator) validateMessages(
	file *ast.ProtoFile,
	messageTypes map[string]struct{},
) {

	messageNames := make(map[string]struct{})

	for _, msg := range file.Messages {
		if msg == nil {
			v.errors = append(
				v.errors,
				"message is required",
			)
			continue
		}

		if msg.Name == "" {
			v.errors = append(
				v.errors,
				"message name is required",
			)
			continue
		}

		if _, exists := messageNames[msg.Name]; exists {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"duplicate message name: %s",
					msg.Name,
				),
			)
		}

		messageNames[msg.Name] = struct{}{}

		v.validateFields(msg, messageTypes)
	}
}

func (v *Validator) validateFields(
	msg *ast.Message,
	messageTypes map[string]struct{},
) {

	fieldNames := make(map[string]struct{})
	fieldNumbers := make(map[int]struct{})

	for _, field := range msg.Fields {
		if field == nil {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"field is required in message %s",
					msg.Name,
				),
			)
			continue
		}

		if field.Name == "" {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"field name is required in message %s",
					msg.Name,
				),
			)
		}

		if field.Type == "" {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"field type is required in message %s",
					msg.Name,
				),
			)
			continue
		}

		if field.Number <= 0 {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"invalid field number in message %s",
					msg.Name,
				),
			)
		}

		if field.Name != "" {
			if _, exists := fieldNames[field.Name]; exists {
				v.errors = append(
					v.errors,
					fmt.Sprintf(
						"duplicate field name '%s' in message %s",
						field.Name,
						msg.Name,
					),
				)
			}

			fieldNames[field.Name] = struct{}{}
		}

		if field.Number > 0 {
			if _, exists := fieldNumbers[field.Number]; exists {
				v.errors = append(
					v.errors,
					fmt.Sprintf(
						"duplicate field number %d in message %s",
						field.Number,
						msg.Name,
					),
				)
			}

			fieldNumbers[field.Number] = struct{}{}
		}

		_, primitive := primitiveTypes[field.Type]
		_, message := messageTypes[field.Type]

		if !primitive && !message {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"unknown type '%s' in message %s",
					field.Type,
					msg.Name,
				),
			)
		}
	}
}

func (v *Validator) validateServices(
	file *ast.ProtoFile,
	messageTypes map[string]struct{},
) {

	serviceNames := make(map[string]struct{})

	for _, svc := range file.Services {
		if svc == nil {
			v.errors = append(
				v.errors,
				"service is required",
			)
			continue
		}

		if svc.Name == "" {
			v.errors = append(
				v.errors,
				"service name is required",
			)
			continue
		}

		if _, exists := serviceNames[svc.Name]; exists {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"duplicate service name: %s",
					svc.Name,
				),
			)
		}

		serviceNames[svc.Name] = struct{}{}

		v.validateRPCs(svc, messageTypes)
	}
}

func (v *Validator) validateRPCs(
	svc *ast.Service,
	messageTypes map[string]struct{},
) {

	rpcNames := make(map[string]struct{})

	for _, rpc := range svc.RPC {
		if rpc == nil {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"rpc is required in service %s",
					svc.Name,
				),
			)
			continue
		}

		if rpc.Name == "" {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"rpc name is required in service %s",
					svc.Name,
				),
			)
		}

		if rpc.Name != "" {
			if _, exists := rpcNames[rpc.Name]; exists {
				v.errors = append(
					v.errors,
					fmt.Sprintf(
						"duplicate rpc name '%s' in service %s",
						rpc.Name,
						svc.Name,
					),
				)
			}

			rpcNames[rpc.Name] = struct{}{}
		}

		if rpc.RequestType == "" {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"request type is required in rpc %s",
					rpc.Name,
				),
			)
		} else if _, exists := messageTypes[rpc.RequestType]; !exists {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"unknown request type '%s' in rpc %s",
					rpc.RequestType,
					rpc.Name,
				),
			)
		}

		if rpc.ResponseType == "" {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"response type is required in rpc %s",
					rpc.Name,
				),
			)
		} else if _, exists := messageTypes[rpc.ResponseType]; !exists {
			v.errors = append(
				v.errors,
				fmt.Sprintf(
					"unknown response type '%s' in rpc %s",
					rpc.ResponseType,
					rpc.Name,
				),
			)
		}
	}
}
