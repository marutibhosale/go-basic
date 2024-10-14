// == > < =>      like these operatores
package main

import "fmt"

func main() {
	// comparision
	var a int = 2
	var b int = 3
	fmt.Println(a == b)
	// arthimatic ++ --
	var x, y string = "foo", "bar"
	fmt.Println(x + y)
	fmt.Println(b % a) // gives reminder
	a++                // increment same way --
	fmt.Println(a)

	// logical    && || !
	fmt.Println((a > 10) && (b > 10))
	fmt.Println(!(a > b)) // not statement

	//assigemtn operatores  = += -= *= /= %=
	a += b // like x = x+y
	println(a)

	// bitwise operators & | ^(XOR) << >>     # its binary bitwase operations
	c := a & b
	println(c)

	// if else
	x = "bar"
	x = "test"
	if x == "foo" {
		println(x)
	} else if x == "bar" { // else if
		println("else if")
	} else {
		println("else") // else
	}

	// switch statements

	switch a {
	case 4:
		println("case 4")
	case 5:
		println("case 5")
	default:
		println("default")
	}

	switch { // switch with condition statements
	case a == b:
		println("case 4")
	case a+b < 60:
		println("case 5")
	default:
		println("default")
	}
	// for loop
	for t := 1; t <= 5; t++ {
		fmt.Println(t * t)
	}

	for t := 1; t <= 5; t++ {
		if t == 3 {
			break // stop the loop and continue is skip current iternation only
		}
		fmt.Println(t * t)
	}

	for t := 1; t <= 5; t++ {
		if t == 3 {
			continue // ontinue is skip current iternation only
		}
		fmt.Println(t * t)
	}

}
