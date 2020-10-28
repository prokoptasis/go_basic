package main

import "fmt"

func main() {
	textLine1 := `첫번째 줄입니다.\n
두번째 줄입니다.\n
	세번째 줄입니다`

	fmt.Println(textLine1)

	textLine2 := "첫번째 줄입니다.\n두번째 줄입니다.\n세번째 줄입니다"

	fmt.Println(textLine2)

	textLine3 := "첫번째 줄입니다.\n" + "두번째 줄입니다.\n" + "세번째 줄입니다"

	fmt.Println(textLine3)

	var a int = 100
	var b uint = uint(a)
	var c float32 = float32(a)

	println(a, b, c)

	str := "ABCDEFG"
	byt := []byte(str)
	stg := string(byt)
	println(str, byt, byt[0], byt[1], stg)
}
