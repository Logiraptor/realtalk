package main

import (
	"reflect"
	"testing"
)

type TestRTObj struct{}

func (TestRTObj) Nil() Object {
	return EmptyObjectVal
}

func (TestRTObj) First(a, b Object) Object {
	return a
}

func (TestRTObj) Second(a, b Object) Object {
	return b
}

func (TestRTObj) CountArgs(one Object, args ...Object) Object {
	return NewSmolInt(len(args) + 1)
}

func TestNativeObject_Send(t *testing.T) {
	type args struct {
		name string
		args []Object
		want Object
	}
	tests := []struct {
		name        string
		nativeValue interface{}
		args        []args
	}{
		{
			name:        "Can redirect to simple methods",
			nativeValue: TestRTObj{},
			args: []args{
				{name: "Nil", want: EmptyObjectVal},
				{name: "First", args: []Object{NewSmolInt(3), NewSmolInt(5)}, want: NewSmolInt(3)},
				{name: "Second", args: []Object{NewSmolInt(3), NewSmolInt(5)}, want: NewSmolInt(5)},
				{name: "CountArgs", args: []Object{NewSmolInt(3), NewSmolInt(5), NewSmolInt(7)}, want: NewSmolInt(3)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNativeObject(tt.nativeValue)
			for _, args := range tt.args {
				if got := n.Send(args.name, args.args...); !reflect.DeepEqual(got, args.want) {
					t.Errorf("NativeObject.Send() = %v, want %v", got, args.want)
				}
			}
		})
	}
}
