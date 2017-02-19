package main

// NilType is a type of object with nothing in it
type NilType struct{}

// RespondTo implements the Object interface
func (n NilType) RespondTo(incoming Message) Object {
	switch incoming.Token {
	case "to_s":
		return StringType("nil")
	case "nil?":
		return True
	}
	return NilObject
}

// NilObject is the instance of NilType used in RealTalk programs
var NilObject = NilType{}

var _ Object = NilType{}
var _ Object = NilObject
