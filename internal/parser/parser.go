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
	if !p.expectPeek(token.ASSIGN) {
		return
	}

	if !p.expectPeek(token.STRING) {
		return
	}
	file.Syntax = p.curToken.Literal

	if !p.expectPeek(token.SEMICOLON) {
		return
	}

}

func (p *Parser) parsePackage(file *ast.ProtoFile) {
	if !p.expectPeek(token.IDENT) {
		return
	}
	pkg := p.curToken.Literal

	for p.peekTokenIs(token.DOT) {
		if !p.expectPeek(token.DOT) {
			return
		}
		if !p.expectPeek(token.IDENT) {
			return
		}
		pkg += "." + p.curToken.Literal
	}
	file.Package = pkg

	if !p.expectPeek(token.SEMICOLON) {
		return
	}

}

func (p *Parser) parseMessage() *ast.Message {
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	msg := &ast.Message{
		Name: p.curToken.Literal,
	}
	return msg
}

func (p *Parser) parseService() *ast.Service {
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	svc := &ast.Service{
		Name: p.curToken.Literal,
	}
	return svc
}
