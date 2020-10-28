package main

import "fmt"

func main() {
	// Array
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	println("arr1[0] : ", arr1[0], "arr1[1] : ", arr1[1], "arr1[2] : ", arr1[2])

	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9}

	println("arr2[0] : ", arr2[0], "arr2[1] : ", arr2[1], "arr2[2] : ", arr2[2])
	println("arr3[0] : ", arr3[0], "arr3[1] : ", arr3[1], "arr3[2] : ", arr3[2])

	// multi array 1
	var arr4 [3][4][5]int
	arr4[0][1][2] = 10
	println("arr4[0][1][2] : ", arr4[0][1][2])

	// multi array 2
	var arr5 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	println("arr5[1][2] : ", arr5[1][2])

	// Slice
	var slc1 []int
	slc1 = []int{1, 2, 3}
	slc1[1] = 10
	fmt.Println("slc1 : ", slc1)

	// Slice with Make
	slc2 := make([]int, 5, 10)
	println(len(slc2), cap(slc2))
	fmt.Println(slc2)

	// Nil Slice
	var slc3 []int

	if slc3 == nil {
		println("Nil Slice")
	}
	println(len(slc3), cap(slc3))
	fmt.Println(slc3)

	// Sub Slice
	slc4 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("[0:1] : ", slc4[0:1], ",[2:5] : ", slc4[2:5], ",[1:] : ", slc4[1:], ",[:5] : ", slc4[:5])

	// Slice append / copy
	slc5 := []int{0, 1}

	// append one
	slc5 = append(slc5, 2)

	// append multiple
	slc5 = append(slc5, 3, 4, 5)

	fmt.Println(slc5)

	// Underlying array
	slc6 := make([]int, 0, 3)

	for i := 1; i < 15; i++ {
		slc6 = append(slc6, i)
		fmt.Println("slc6 len, cap : ", len(slc6), cap(slc6))
	}

	fmt.Println(slc6)

	// append slice & ellipsis(...)
	slc7 := []int{1, 2, 3}
	slc8 := []int{4, 5, 6}
	slc7 = append(slc7, slc8...)
	fmt.Println(slc7)

	// slice copy
	slc9 := []int{0, 1, 2}
	slc10 := make([]int, len(slc9), cap(slc9)*2)
	copy(slc10, slc9)
	fmt.Println(slc10)
	println(len(slc10), cap(slc10))
}
