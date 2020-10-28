package main

import (
	"fmt"
	"os"
)

func main() {
	// // Error 처리
	// f, err := os.Open("C:\\temp\\1.txt")
	// if err != nil {
	// 	println("Error!!!")
	// 	log.Fatal(err.Error())
	// }
	// println(f.Name())

	// // defer
	// f, err := os.Open("1.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// bytes := make([]byte, 1024)
	// f.Read(bytes)
	// println(len(bytes))

	// panic
	openFile("Invalid.txt")
	println("Done!!!")
}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open Error", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		println("Panic Herer!!!")
		panic(err)
	}

	defer f.Close()
}
