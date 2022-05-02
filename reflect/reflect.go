package main

import (
	"fmt"
	"reflect"
)

type serialNum interface {
	serial() string
}

type student struct {
	name string
	id   int
}

func (s student) serial() string {
	str := fmt.Sprintf("%d%s", s.id, s.name)

	return str
}

func displaySN(s serialNum) error {

	fmt.Println(s.serial())

	sv := reflect.ValueOf(s)

	fmt.Println(reflect.TypeOf(s), reflect.ValueOf(s))
	fmt.Println(sv.Kind(), sv.Type(), sv.Interface())

	return nil
}

func main() {
	s := student{"frank", 122}
	_ = displaySN(s)
}
