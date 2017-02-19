package interp

import (
	"fmt"
)

type Expression interface {
	_expr()
}

var _ Expression = LiteralExpression{}
var _ Expression = SendExpression{}

func (LiteralExpression) _expr() {}
func (SendExpression) _expr()    {}

type LiteralExpression struct {
	Value Object
}

type SendExpression struct {
	Receiver Expression
	Name     string
	Args     []Expression
}

func InterpExpression(expr Expression) Object {
	switch v := (expr).(type) {
	case LiteralExpression:
		return v.Value
	case SendExpression:
		args := make([]Object, len(v.Args))
		for i, a := range v.Args {
			args[i] = InterpExpression(a)
		}
		receiver := InterpExpression(v.Receiver)
		return receiver.Send(v.Name, args...)
	}
	panic(fmt.Sprintf("Cannot interpret expression of type %T", expr))
}
