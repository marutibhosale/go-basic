package main

import "fmt"

// pointeres are varialbes which stores the memory address of another variable
func main() {
	i := 10
	j := &i
	fmt.Printf("%T %v \n", &i, &i)       // & shows address of the variable
	fmt.Printf("%T %v \n", *(&i), *(&i)) // * gives value at address
	fmt.Println(*(j))
	fmt.Println("address valune is", *(j))

	//var prt_i *int // pointter initilization empty

	//var prt_in *int = &i // variable ptr_in store value of address of variable i
	var ptr = &i // this also work
	//ptr_t := &i          // short hand decalration

	*ptr = 110 // here changed of i using defereing pointer like
	// i = *(&i)
	fmt.Println(i)
	// func mofigy (s *string) {     // we can pass reference by varialbe
	//  *s = world                   // assigened new valune to that addredd
	// }
	// a := "heloow"
	// mofigy(&a)     // passed addredd of variable

	////// Maps and slices are passed by reference by default ///////
}
