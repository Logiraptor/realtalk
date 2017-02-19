package main

var EmptyObjectVal = EmptyObject{}

type EmptyObject struct{}

var _ Object = EmptyObject{}

func (EmptyObject) Send(name string, args ...Object) Object {
	panic("Cannot send messages to the empty object")
}

func (EmptyObject) String() string {
	return "<empty object>"
}
