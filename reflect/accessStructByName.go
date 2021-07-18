package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X MiniVertex
	Y int
}

type MiniVertex struct {
	XX int
	YY int
}

func main() {
	m := MiniVertex{1, 2}
	v := Vertex{m, 2}
	Xval := v.getField("X")
	newXval := Xval.Interface().(MiniVertex)
	XXval := newXval.getField("XX")
	fmt.Println(XXval)
}

func (v *Vertex) getField(field string) reflect.Value {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	//return int(f.Int())
	return f
}

func (v *MiniVertex) getField(field string) reflect.Value {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	//return int(f.Int())
	return f
}
