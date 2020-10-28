package main

import "fmt"

func main() {
	// map 1
	var map01 map[int]string

	map01 = make(map[int]string)
	map01[901] = "Apple"
	map01[134] = "Grape"
	map01[777] = "Tomato"

	str := map01[134]
	println("134 :", str)
	str = map01[999]
	println("999 :", str)
	str = map01[777]
	println("777 :", str)
	delete(map01, 777)
	println("777 :", str)

	// map key check
	map02 := map[string]string{
		"GOO": "Google",
		"MSF": "Microsoft",
		"FBK": "Facebook",
		"AMZ": "Amazon",
	}

	val, exists := map02["MSF"]
	if !exists {
		println("No MSF MAP02")
	}

	println("Exists :", val, exists)

	// map for loop
	map03 := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	for key, val := range map03 {
		fmt.Println(key, val)
	}
}
