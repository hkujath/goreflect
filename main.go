package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyType string
	var (
		myInt    int    = 12
		myString string = "Hello"
		myBool   bool   = true
		myType   MyType = "MyType"
	)
	refl(myInt)
	refl(myString)
	refl(myBool)
	refl(myType)

	// Changing a value by using reflection
	refl2(myInt)
	fmt.Println(myInt)
	refl2(&myInt)
	fmt.Println(myInt)

	myStruct := struct {
		Field1 string
		Field2 string
	}{
		Field1: "Value Field1",
		Field2: "Value Field2",
	}
	refl3(myStruct)
}

func refl(i interface{}) {
	fmt.Println("Type:", reflect.TypeOf(i))
	val := reflect.ValueOf(i)
	fmt.Println("Value:", val, val.Kind())
}

func refl2(i interface{}) {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.CanSet() {
		val.SetInt(99)
	}
}

func refl3(i interface{}) {
	val := reflect.ValueOf(i)
	type1 := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := type1.Field(i)
		fmt.Println("Name:", field.Name)
		fmt.Println("Value:", val.Field(i))
	}
}
