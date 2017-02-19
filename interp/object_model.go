package interp

import (
	"fmt"
)

var rtBasicObject = NewRTClass("BasicObject", nil)
var rtObject = NewRTClass("Object", rtBasicObject)
var rtClass = NewRTClass("Class", rtObject)

type Object interface {
	Class() Object
	Send(string, ...Object) Object
}

type Class interface {
	ResolveMethod(name string) (Method, bool)
	Name() string
	Object
}

type Method interface {
	Call(self Object, args ...Object) Object
}

type VariableAssignment struct {
	vars map[string]*Object
}

func NewVariableAssignment() *VariableAssignment {
	return &VariableAssignment{
		vars: make(map[string]*Object),
	}
}

func (v *VariableAssignment) Lookup(name string) *Object {
	return v.vars[name]
}

func (v *VariableAssignment) Assign(name string, value *Object) *Object {
	v.vars[name] = value
	return value
}

type MethodSet struct {
	KnownMethods map[string]Method
}

func NewMethodSet() *MethodSet {
	return &MethodSet{
		KnownMethods: make(map[string]Method),
	}
}

func (ms *MethodSet) AddMethod(name string, body Expression) {
	ms.KnownMethods[name] = &RTMethod{
		Body: body,
	}
}

type RTMethod struct {
	Body Expression
}

func (m *RTMethod) Call(self Object, args ...Object) Object {
	return InterpExpression(m.Body)
}

type RTClass struct {
	Methods    *MethodSet
	SuperClass Class
	name       string
}

func NewRTClass(name string, superClass Class) *RTClass {
	return &RTClass{
		name:       name,
		Methods:    NewMethodSet(),
		SuperClass: superClass,
	}
}

func (c *RTClass) Class() Object {
	return rtClass
}

func (c *RTClass) ResolveMethod(name string) (Method, bool) {
	if c == nil {
		return nil, false
	}
	if m, ok := c.Methods.KnownMethods[name]; ok {
		return m, true
	}
	return c.SuperClass.ResolveMethod(name)
}

func (c *RTClass) Send(name string, args ...Object) Object {
	return nil
}

func (c *RTClass) Name() string {
	return c.name
}

type RTObject struct {
	class         Class
	InternalState *VariableAssignment
}

func NewRTObject(class Class) Object {
	return &RTObject{
		class:         class,
		InternalState: NewVariableAssignment(),
	}
}

func (c *RTObject) Class() Object {
	return c.class
}

func (o *RTObject) Send(name string, args ...Object) Object {
	if m, ok := o.class.ResolveMethod(name); ok {
		return m.Call(o, args...)
	}
	return nil
}

func (o *RTObject) String() string {
	return fmt.Sprintf("<%s %p>", o.class.Name(), o)
}
