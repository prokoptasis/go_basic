package main

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	println(sum)

	n := 2

	for n < 100 {
		n = n * n
	}

	println(n)

	names := []string{"abc", "def", "geh"}

	for index, name := range names {
		println(index, name)
	}

	b := 0

L1:
	for {
		if b == 0 {
			break L1
		}
	}

	println("L1", b)

	var a = 1
	for a < 15 {
		if a == 5 {
			a += 4
			continue
		}
		a++
		println("In", a)
		if a > 9 {
			println("Here", a)
			break
		}
	}
	if a == 11 {
		goto END
	}
	println(a)

END:
	println("End", a)
}
