package main

import "fmt"

// struct
type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	dic := dict{}
	dic.data = map[int]string{}
	return &dic
}

func main() {
	// person struct
	per1 := person{}

	per1.name = "Lee"
	per1.age = 18

	fmt.Println(per1)

	// person struct 2
	var per2 person
	per2 = person{"Bob", 20}
	per3 := person{name: "Sean", age: 30}
	fmt.Println(per2)
	fmt.Println(per3)

	// person struct 3
	per4 := new(person)
	per4.name = "Kim"

	fmt.Println(per4)

	// constructor
	dic1 := newDict()
	dic1.data[1] = "A"

	fmt.Println(dic1)
}
