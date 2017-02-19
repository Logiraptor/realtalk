package interp

type PrimitiveClass struct {
	name string
}

func (p PrimitiveClass) Name() string {
	return p.name
}

func (PrimitiveClass) Class() Object {
	return rtClass
}

func (PrimitiveClass) Send(name string, args ...Object) Object {
	return nil
}

type RTInt int

var _ Object = RTInt(0)

func (RTInt) Send(name string, args ...Object) Object {
	return nil
}

func (RTInt) Class() Object {
	return PrimitiveClass{name: "Integer"}
}
