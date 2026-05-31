package ast

type ProtoFile struct {
	Syntax  string
	Package string

	Messages []*Message
	Services []*Service
}

type Message struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Type   string
	Name   string
	Number int
}

type Service struct {
	Name string
	RPC  []*RPC
}

type RPC struct {
	Name         string
	RequestType  string
	ResponseType string
}
