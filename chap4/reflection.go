package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Lucas", "Jeon"}}

	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String()) // String value: <main.Record Value>

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType) // i Type: main.Record
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		// Field1  with type: string       and value _String value_
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}

		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
