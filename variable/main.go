package main

import "fmt"

func main() {
    // string
    var name string
    name = "Faiyaz Ahmed"

    firstname := "Faiyaz"
    var lastname = "Ahmed"

    fmt.Println(name, firstname, lastname)

    // number
    var num = 1
    fmt.Println("Number:", num)

    var number int = 34
    fmt.Println("Number:", number)

    num1 := 15
    num1 = 30
    fmt.Println("Number:", num1)

    // bool
    isLogin := false
    fmt.Println("isLogin:", isLogin)

    var isOk bool
    isOk = true
    fmt.Println("isOk:", isOk)

	fmt.Printf("%T",   num1)
}
