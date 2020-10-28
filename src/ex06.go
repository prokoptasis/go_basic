package main

func main() {
	msg := "Hello"

	// Function
	say(msg)
	println("2", msg)

	// Pass By Reference
	sayC(&msg)
	println("4", msg)

	// Variadic Function
	sayV("this", "is", "a", "pen")
	sayV("and", "pineapple")

	// Return Value
	total := sum(1, 7, 3, 5, 9)
	println(total)

	// Return Multiple Value
	count, total := sumC(1, 7, 3, 5, 9)
	println(count, total)

	// Named Return Parameter
	count, total = sumN(1, 7, 3, 5, 9)
	println("Named Return Param", count, total)

	// Anonymous Function
	sumA := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sumA(1, 2, 3, 4, 5)
	println("result :", result)

	// add anonymouse function
	add := func(i int, j int) int {
		return i + j
	}

	// send anonumous function to other function
	r1 := calc(add, 10, 20)
	println("r1 :", r1)

	// directly
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println("r2 :", r2)

	// prototype function
	r3 := calcP(add, 10, 25)
	println("r3 :", r3)

	// closure
	next1 := nextValue()
	println("next1", next1())
	println("next1", next1())
	println("next1", next1())

	next2 := nextValue()
	println("next2", next2())
	println("next2", next2())
	println("next1", next1())
}

func say(msg string) {
	println("1", msg)
}

func sayC(msg *string) {
	println("3", *msg)
	*msg = "World"
}

func sayV(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sumC(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sumN(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// prototype
type calculator func(int, int) int

// use prototype function
func calcP(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// closure
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
