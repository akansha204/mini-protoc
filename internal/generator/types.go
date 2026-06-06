package generator

var typeMap = map[string]string{
	"string": "string",
	"int32":  "int32",
	"int64":  "int64",
	"bool":   "bool",
	"float":  "float32",
	"double": "float64",
}

func goType(protoType string) string {
	if goType, ok := typeMap[protoType]; ok {
		return goType
	}
	return protoType
}
