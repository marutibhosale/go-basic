package main

import "fmt"

func main() {
	//arry fo fixed length ocne decalre
	var grades [5]int // array declaration
	fmt.Println(grades)
	var grade [3]int = [3]int{1, 2, 3}
	fmt.Println(grade)
	grad := [3]int{5, 2, 3}
	fmt.Println(grad)
	testt := [...]int{1, 2, 3, 4, 5} // [...]  no need to define array length
	fmt.Println(testt)
	fmt.Println(len(testt)) // len
	fmt.Println(testt[1])   // indexing
	testt[1] = 10           // assine new valune
	fmt.Println(testt)

	for i := 0; i < len(testt); i++ {
		fmt.Println(testt[i])
	}
	for index, element := range testt {
		fmt.Println(index, "=>", element)
	}

	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}} // multi D array
	fmt.Println(arr[1][0])

	// slice  array[start_index : end_index]
	//
	fmt.Println(testt[1:3])  // 3rd not include
	fmt.Println((testt[:3])) // start 0 index
	fmt.Println(testt[1:])   // last index
	fmt.Println(testt[:])    // whole array sliced
	// we can take sub slic also

	slic := make([]int, 5, 8) // cap of slic
	fmt.Println(slic)
	println(len(slic))
	println(cap(slic))

	slic_1 := testt[:4]
	fmt.Println(testt)
	fmt.Println(slic_1)
	slic_1[1] = 9999 // not only slic index value change but also change array value also
	fmt.Println(testt)
	fmt.Println(slic_1)

	slic_1 = append(slic_1, 34, 45, 54) // we can append one slic to another
	fmt.Println(slic_1)
	fmt.Println(testt)
	// we can delete element in slic by creating two slices by exclucing removing element and append two slices together
	slic2 := make([]int, 3)
	copy(slic2, slic_1) // copying from one slic to another
	fmt.Println(slic2)

	// Maps
	//var test map[string]int // string is key data type and int is value data type
	dict := map[string]string{"in": "hindi", "mh": "marathi"}
	fmt.Println(dict["in"])
	dict["in"] = "HN"
	fmt.Println(dict["in"])
	codes := make(map[string]int)
	fmt.Println(codes)
	value, found := dict["mh"]
	fmt.Println(value, found) // found is boolen
	delete(dict, "in")
	fmt.Println(dict)
	dict1 := map[string]string{"in": "hindi", "mh": "marathi"}

	for key, value := range dict1 { // iterate over key value
		fmt.Println(key, ":>", value)
	}

}
