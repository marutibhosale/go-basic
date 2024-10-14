package main

import "fmt"

func add(a int, b int) int { // last int is reture type we can use multiple returen tyoe like (int, string) or named
	// named type like (sum int, diff int)
	sum := a + b
	return sum
}

// verdic function

func ver(num ...int) int { // ... we can pass as many arguments for num int type
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

// recursive function
func factorial(n int) int {
	if n == 0 {  
		return 1
	}
	return n * factorial(n-1)

}

func main() {
	addition := add(2, 3)
	fmt.Println(addition)
	fmt.Println(ver(1, 2, 3, 4, 5))
	// recursive
	n := 5
	result := factorial(n)
	fmt.Println("Factorial fo", n, "is : ", result)
	// anonymous function
	x := func(a int, b int) int {
		return a * b
	}(10, 20)
	fmt.Printf("%T \n", x)
	//fmt.Println(x(10,20))
	fmt.Println(x)

	// high order function

	//defer

}
