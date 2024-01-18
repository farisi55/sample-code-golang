package main //shortcut pkgm

import (
	"fmt"
	"reflect"
)

//shortcut tys
type student struct {
	Nama string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem() //mengambil object reflect aslinya maka nya keyword Elem
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama Filed :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func main() {
	//var s1 = &student{Nama: "Banu", Grade: 3}
	var s1 = &student{"Banu", 3}
	s1.getPropertyInfo()
}