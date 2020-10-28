package main

func main() {
	var a int = 7
	var b int = 3

	//산술연산자
	var c = (a + b) / 5

	println(a, b, c)

	c++
	println(c)

	// 관계연산자
	if b == c {
		println("b==c")
	} else {
		println("b!=c")
	}

	if a != b {
		println("a!=b")
	} else {
		println("a==b")
	}

	if a >= b {
		println("a>=b")
	} else {
		println("a< b")
	}

	// 논리 연산자
	if a >= b && b == c {
		println("a>=b && b==c")
	} else {
		println("not(a>=b && b==c)")
	}

	if a >= b || b > c {
		println("a>=b || b<>c")
	} else {
		println("not(a>=b || b<>c)")
	}

	// Bitwise 연산자
	var d = (a & b)
	var e = (a & b) << 1

	println(d, e)

	// 할당연산자
	a = 10
	println(a)
	a *= 10
	println(a)
	a >>= 1
	println(a)
	a |= 1
	println(a)

	// 포인터 연산자
	var x int = 10
	var p = &x //x's address

	// x = value
	// &x = x's address
	// p = &x = x's address
	// *p = x's address's value
	// &p = p's address
	println(x, &x, p, *p, &p)

	x++

	println(x, &x, p, *p, &p)
}
