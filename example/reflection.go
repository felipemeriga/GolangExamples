package example

import (
	"fmt"
	"reflect"
)

type Register struct {
	Character Character
}

type Character struct {
	Level int
	Data  Data
}

type Data struct {
	Name string
	Age  int
}

func ReflectionOnCharacter() {
	randomCharacter := &Register{Character: Character{
		Level: 12,
		Data: Data{
			Name: "Felipe",
			Age:  30,
		},
	}}

	InspectStructV(reflect.ValueOf(randomCharacter))
	//rCharacter := reflect.ValueOf(randomCharacter)
	//iterateReflect(rCharacter)
}

func InspectStructV(val reflect.Value) {
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		address := "not-addressable"

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()

		}
		if valueField.CanAddr() {
			address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
		}

		if typeField.Name == "Age" {

			valueField.Set(reflect.ValueOf(23))
		}

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n", typeField.Name,
			valueField.Interface(), address, typeField.Type, valueField.Kind())

		if valueField.Kind() == reflect.Struct {
			InspectStructV(valueField)
		}
	}
}

func iterateReflect(v reflect.Value) {
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			f := reflect.Indirect(v.Field(i))

			if f.Type().Name() == "Data" {
				fmt.Println("found")
			}
			iterateReflect(f)
		}
	} else {
		printReflect(v)
	}

}

func printReflect(v reflect.Value) {
	fmt.Println(fmt.Sprintf("value: %v", v))
}
