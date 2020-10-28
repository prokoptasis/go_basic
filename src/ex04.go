package main

func main() {
	var a = 1

	if a == 1 {
		println("One")
	} else if a == 2 {
		println("Two")
	} else {
		println("Other")
	}

	// Optional Statement
	if b := 1; b < 10 {
		println(b)
	}

	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	// Expression을 사용한 경우
	switch x := category << 2; x - 1 {
	//...
	}

	var score = 80

	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}
