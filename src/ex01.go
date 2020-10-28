package main

func main() {
	// 선언 및 할당
	var a int
	var b float32 = 11.

	println(a, b)

	// 선언 및 할당
	a = 10
	b = 12.0

	println(a, b)

	// 상수
	const c int = 10
	const d string = "Hello"

	println(c, d)

	// 상수 나열
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	println(Visa, Master, Amex)

	// 상수 나열
	const (
		Apple = iota
		Grape
		Orage
	)

	println(Apple, Grape, Orage)

}
