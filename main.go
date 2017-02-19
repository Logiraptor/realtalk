package main

type Object interface {
	Send(name string, args ...Object) Object
}

func main() {

}
