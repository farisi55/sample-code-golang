package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	var name = "Banu salman"
	var reflectValueString = reflect.ValueOf(name)

	fmt.Println("tipe variable number :", reflectValue.Type())
	fmt.Println("tipe variable name :", reflectValueString.Type())

	/*
		untuk kasus ini. Interface() hanya diperlukan utk ditampilkan di output
		jadi lewat method ini segala jenis nilai bisa diakses dengan mudah
	*/

	fmt.Println("nilai variable number interfacenya :", reflectValue.Interface())
	var nilai = reflectValue.Interface().(int) // casting to int
	fmt.Println("nilai asli (int) :", nilai)

	/*
		line 23, sebenarnya mengembalikan nilai interface kosong atau interface{}
		dengan cara mengcasting interface kosong
	*/

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable :", reflectValue.Int())
	}
}