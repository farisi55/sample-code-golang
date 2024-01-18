package main

import ("fmt")

func Ups(i int) interface{}{
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else {
		return "ups"
	}
}

func Hello(name string) interface{}{
	if name == "" {
		return "apa"
	}else if name == "budi"{
		return false
	}else {
		return 123
	}
}

func Banyak(name string, i int) interface{}{
	//return fmt.Println("name", name, "usia", i) 

	if name == ""  && i == 0 {
		return "apa"
	}else if name == "budi" && i == 0 {
		return false
	}else {
		return 321
	}
}

func main() {
	var data interface{} = Ups(3)

	var kata interface{} = Hello("tono")

	var test interface{} = Banyak("budi", 17)

	fmt.Println(data)
	fmt.Println("hai", kata)
	fmt.Println(test)
	
}