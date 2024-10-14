package main

// struct is user define data type

import "fmt"

type student struct {
	name   string
	rollno int
	//marks  []int
	//grades map[string]int
}

func test(s student) { // like we can psss struct to function

}

func (c *Circle) calArea() { // its called method here we passed struct before function
	c.area = 3.14 * c.redius * c.redius
}

func (c *Circle) printredis() {
	fmt.Printf("Redius is %+v \n", c.redius)
}

type Circle struct {
	redius float64
	area   float64
}

// interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct { // struct
	side float64
}

func (s square) area() float64 { // method 1
	return s.side * s.side
}

func (s square) perimeter() float64 { // method 2
	return s.side * 4
}

type react struct { // struct
	length  float64
	breadth float64
}

func (s react) area() float64 { // method 1
	return s.length * s.breadth
}

func (s react) perimeter() float64 { // method 2
	return 2*s.length + 2*s.breadth
}

func printdata(s shape) { // print for two methods area and permiter
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func main() {
	var s student           // student is user define data type like int string
	fmt.Printf("%+v \n", s) // +v
	st := student{"maruti", 12}
	fmt.Printf("%+v \n", st)
	st.name = "joe"
	st.rollno = 43
	fmt.Printf("%+v \n", st)
	//stu := student{"maruti", 12}
	//ste := student{"maruti", 12} // compare strct like if stu == ste and !=
	//test(stu)     //  calling function by passing struct

	c := Circle{redius: 5}
	c.calArea()    // struct.functon possible for method only
	c.printredis() // if multiple method call then it call method set
	fmt.Printf("%+v \n", c)

	// interface

	sq := square{side: 4}
	re := react{length: 3, breadth: 4}
	printdata(sq) // p
	printdata(re)
}
