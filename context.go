package main

import (
	"fmt"

	"github.com/Logiraptor/realtalk/ast"
)

// Context represents an execution environment
type Context struct {
	Bindings Bindings
}

// Bindings are a set of variable bindings
type Bindings interface {
	Set(string, Object)
	Lookup(string) Object
}

// Evaluate evaluates a program in the context.
func (c Context) Evaluate(p ast.Program) {
	for _, st := range p.Statements {
		c.EvaluateExpression(st)
	}
}

// EvaluateExpression evaluates the expression in the context.
func (c Context) EvaluateExpression(e ast.Expression) Object {
	switch v := e.(type) {
	case ast.Assignment:
		rhs := c.EvaluateExpression(e)
		c.Bindings.Set(v.Identifier.Token, rhs)
		return rhs
	case ast.SendExpression:
		receiver := c.EvaluateExpression(v.Receiver)
		message := c.EvaluateExpression(v.Message)
		if m, ok := message.(Message); ok {
			return receiver.RespondTo(m)
		}
		panic(fmt.Sprintf("object cannot respond to non-message: %T", message))
	case ast.Identifier:
		return c.Bindings.Lookup(v.Token)
	default:
		panic(fmt.Sprintf("cannot evaluate non-expression: %T", e))
	}
}
