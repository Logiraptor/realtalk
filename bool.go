package main

var _ Object = BoolType(false)
var _ Object = False
var _ Object = True

// BoolType is the primitive bool
type BoolType bool

// False is the global false value
var False = BoolType(false)

// True is the global true value
var True = BoolType(true)

// RespondTo implements the Object interface
func (b BoolType) RespondTo(incoming Message) Object {
	switch incoming.Token {
	case "to_s":
		if b {
			return StringType("true")
		}
		return StringType("false")
	}
	return NilObject
}
