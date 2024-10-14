package main

import (
	"fmt"
	"strconv"
	//"reflect"
)

var var_global string = "test" // global var
func main() {
	// fmt.Println("hellow")
	// fmt.Print("hellow")
	// var test string = "world"
	// fmt.Println("welcome to new", test)
	// fmt.Println(test)
	// var name = "maruti"
	// fmt.Printf("how are you %v \n", name) // v get name value like %d , %c
	// var city string                       // initilizaion first then declaration
	// city = "pune"
	// fmt.Println(city)
	// town := "kavate"  // no need to mention var keyword and data type name like string
	// town = "makankal" // can change name of variable
	// fmt.Println(town)

	// var surname string
	// fmt.Print("enter your surname: ")
	// fmt.Scanf("%s", &surname) // input function
	// fmt.Println(surname)
	// var username string
	// var age int
	// fmt.Print("enter name and age: ")
	// fmt.Scanf("%s  %d", &username, &age) // multiple input
	// fmt.Println(username, age)
	// var a string
	// var b int
	// fmt.Print("enter valune a and b: ")
	// count, err := fmt.Scanf("%s %d", &a, &b)
	// fmt.Println("count", count)
	// fmt.Println("err", err)
	// fmt.Println("a ", a)
	// fmt.Println("b ", b)

	// fmt.Printf("varible %v and its type %T \n", a, a)      // %v for value and %T for type varaible type
	// fmt.Printf("varaible type %T \n", false)               // %T for get type of variable
	// fmt.Printf("variable b type %v \n", reflect.TypeOf(b)) // we can use reflect.TypeOf also

	// converting between data types

	var i int = 90
	var f float64 = float64(i)

	fmt.Printf("%.2f \n", f)

	fmt.Printf("%v %T \n", i, strconv.Itoa(i))

	// constant

	const a = 10
	fmt.Printf("%v %T \n", a, a)
}

// %v, %+v and %T in printf , for scanf we can use %s and %d
