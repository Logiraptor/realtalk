package main

import (
	"fmt"

	"github.com/Logiraptor/realtalk/interp"
)

func printObject(obj Object) {
	fmt.Println(formatObject(obj))
}

func formatObject(obj Object) string {
	return string(obj.RespondTo(Message{
		Token: "to_s",
	}).(StringType))
}

func main() {
	object := interp.NewRTClass("Object", nil)
	object.Methods.AddMethod("print", interp.LiteralExpression{
		Value: interp.RTInt(5),
	})

	instance := interp.NewRTObject(object)
	fmt.Println(instance.Send("print"))
}
