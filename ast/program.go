package ast

import "github.com/Logiraptor/realtalk/token"

// Program represents an executable realtalk program
type Program struct {
	Statements []Expression
}

type Expression interface {
}

// var _ Expression = Object(nil)
var _ Expression = Assignment{}
var _ Expression = SendExpression{}
var _ Expression = Identifier{}
var _ Expression = Int(0)

type Assignment struct {
	Identifier Identifier
	Value      Expression
}

type SendExpression struct {
	Receiver Expression
	Message  Expression
}

type Identifier struct {
	Token string
}

type Int int

type Attrib interface{}

func NewID(a Attrib) Identifier {
	return Identifier{
		Token: string(a.(token.Token).Lit),
	}
}
