package parser

import (
	"fmt"

	"github.com/akansha204/mini-protoc/internal/ast"
	"github.com/akansha204/mini-protoc/internal/lexer"
	"github.com/akansha204/mini-protoc/internal/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		t,
		p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProtoFile() *ast.ProtoFile {
	file := &ast.ProtoFile{}

	for !p.curTokenIs(token.EOF) {

		switch p.curToken.Type {

		case token.SYNTAX:
			p.parseSyntax(file)

		case token.PACKAGE:
			p.parsePackage(file)

		case token.MESSAGE:
			msg := p.parseMessage()
			file.Messages = append(file.Messages, msg)

		case token.SERVICE:
			svc := p.parseService()
			file.Services = append(file.Services, svc)
		}

		p.nextToken()
	}

	return file
}

func (p *Parser) parseSyntax(file *ast.ProtoFile) {
}

func (p *Parser) parsePackage(file *ast.ProtoFile) {
}

func (p *Parser) parseMessage() *ast.Message {
	return nil
}

func (p *Parser) parseService() *ast.Service {
	return nil
}
