package main

var _ Object = StringType("")

// StringType is the primitive string
type StringType string

// RespondTo implements the Object interface
func (s StringType) RespondTo(incoming Message) Object {
	switch incoming.Token {
	case "to_s":
		return s
	}
	return NilObject
}
