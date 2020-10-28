package main

import (
	"fmt"
	"math"
)

// struct
type Rect struct {
	width, height int
}

// method
func (r Rect) area() int {
	return r.width * r.height
}

// method
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

// interface
type Shpe interface {
	area() float64
	peri() float64
}

type RectF struct {
	width, height float64
}

type CircF struct {
	radius float64
}

// interface
func (r RectF) area() float64 { return r.width * r.height }
func (r RectF) peri() float64 { return 2 * (r.width + r.height) }
func (c CircF) area() float64 { return math.Pi * c.radius * c.radius }
func (c CircF) peri() float64 { return 2 * math.Pi * c.radius }

func showArea(shpe ...Shpe) {
	for _, s := range shpe {
		a := s.area()
		println(a)
	}
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
	area = rect.area2()
	println(area)

	rectf := RectF{10., 20.}
	circf := CircF{10}
	showArea(rectf, circf)

	// empty interface
	var type_x interface{}
	type_x = 1
	type_x = "Jake"

	printIt(type_x)

	// Type Assertion
	var type_y interface{} = 1
	var_i := type_y
	var_j := type_y.(int)

	println(var_i) // pointer address
	println(var_j) // 1
}

func printIt(type_v interface{}) {
	fmt.Println(type_v)
}
