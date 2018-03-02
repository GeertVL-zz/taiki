package ast

import "github.com/geertvl/taiki/token"

type Node interface {
	TokenLiteral() string
}

type Diagram interface {
	Node
	diagramNode()
}

type Expression interface {
	Node
	expressionNode()
}

/// Program
type Program struct {
	Diagrams []Diagram
}

func (p *Program) TokenLiteral() string {
	if len(p.Diagrams) > 0 {
		return p.Diagrams[0].TokenLiteral()
	}

	return ""
}

/// SequenceDiagram
type SequenceDiagram struct {
	Token token.Token
	Name *Identifier
	Expressions []Expression
}

func (sd *SequenceDiagram) diagramNode() {}
func (sd *SequenceDiagram) TokenLiteral() string { return sd.Token.Literal }

/// Identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }


