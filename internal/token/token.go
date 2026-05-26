package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	//identifiers and literals
	IDENT  TokenType = "IDENT"
	NUMBER TokenType = "NUMBER"
	STRING TokenType = "STRING"

	//operators
	ASSIGN TokenType = "="

	//delimiters
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	DOT       TokenType = "."

	//keywords
	MESSAGE TokenType = "MESSAGE"
	SERVICE TokenType = "SERVICE"
	RPC     TokenType = "RPC"
	RETURNS TokenType = "RETURNS"
	SYNTAX  TokenType = "SYNTAX"
	PACKAGE TokenType = "PACKAGE"
)

var keywords = map[string]TokenType{
	"message": MESSAGE,
	"service": SERVICE,
	"rpc":     RPC,
	"returns": RETURNS,
	"syntax":  SYNTAX,
	"package": PACKAGE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
