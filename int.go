package main

type SmolInt int

func NewSmolInt(i int) Object {
	return NewNativeObject(SmolInt(i))
}
