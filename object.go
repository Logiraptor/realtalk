package main

import (
	"fmt"
)

// Object is the type of all things in RealTalk
type Object interface {
	RespondTo(Message) Object
}

// Message represents some data that may trigger computation when given to an object
type Message struct {
	Token string
	Data  Object
}

// RespondTo implements the Object interface
func (m Message) RespondTo(incoming Message) Object {
	switch incoming.Token {
	case "token":
		return StringType(m.Token)
	case "data":
		return m.Data
	case "to_s":
		return StringType(fmt.Sprintf("[Message %s %v]", m.Token, m.Data))
	}
	return NilObject
}

var _ Object = Message{}
