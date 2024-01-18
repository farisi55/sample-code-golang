package main

import "fmt"

func main(){
    for counter := 1; counter <= 10; counter++{
        fmt.Println("Perulangan ke", counter)
    }

    slice := []string{"banu", "salman", "farisi", "bogor", "jakarta"}

    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    for i, value := range slice {
        fmt.Println("Index", i, "=", value)
    }

    person := make(map[string]string)
    person["name"] = "Feri"
    person["title"] = "Programmer"

    for key, value := range person {
        fmt.Println(key, "=", value)        
    }

    m := map[string]map[string]string{}
    m["var1"] = map[string]string{}
    m["var1"] ["var2"] = "something"
    fmt.Println(m["var1"]["var2"])


}