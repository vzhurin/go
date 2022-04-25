package main

import (
	"fmt"
	"math"
)

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}

func variable() {
	name := "Basil"
	fmt.Printf("Hello, %v!\n", name)
	fmt.Println("My name len is", len(name))
	fmt.Println(ftoc(70))
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for k := 0; k <= 10; {
		if k%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

		k += 1
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func switch_() {
	j := 7

	switch j {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Any")
	}

	for i := 1; i <= 100; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
		divBoth := div3 && div5

		switch true {
		case divBoth:
			fmt.Println("FizzBuzz")
		case div3:
			fmt.Println("Fizz")
		case div5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func array() {
	a1 := [5]int{1, 9}
	fmt.Println(a1)

	var a2 [5]int
	fmt.Println(a2)
	fmt.Println(len(a2))

	total := 0.0
	for i, value := range a1 {
		total += float64(value)
		fmt.Printf("%v: %v\n", i, value)
	}

	fmt.Println(total / float64(len(a1)))
}

func slice() {
	sl1 := []float64{1, 2, 3, 4, 5}
	fmt.Println(sl1)

	sl2 := make([]float64, 5)
	fmt.Println(sl2)

	// 10 is the length of the underlying array
	sl3 := make([]float64, 5, 10)
	fmt.Println(sl3)

	sl4 := sl3[0:5]
	fmt.Println(sl4)

	sl5 := sl1[:]
	fmt.Println(sl5)

	// Append
	sl6 := []int{1, 2, 3, 4, 5}
	sl7 := append(sl6, 7, 10)
	fmt.Println(sl7)

	// Copy
	sl8 := []int{1, 2, 3, 4, 5}
	sl9 := []int{7, 8, 9, 10, 11, 12, 13, 15}
	copy(sl8, sl9)

	fmt.Println(sl8)

	for i, value := range sl9[2:7] {
		fmt.Println(i, value)
	}
}

func map_() {
	m1 := map[string]int{"b": 2}
	m1["a"] = 1
	fmt.Println(m1)

	// Error
	//var m2 map[string]int
	//m2["a"] = 1
	//fmt.Println(m2)

	m2 := make(map[string]int)
	m2["a"] = 1
	m2["b"] = 2
	fmt.Println(m2)
	fmt.Println(len(m2))

	delete(m2, "a")
	fmt.Println(m2)
	fmt.Println(len(m2))

	value, ok := m2["bla"]
	fmt.Println(value, ok)

	if value, ok := m2["b"]; ok {
		fmt.Println(value, ok)
	}

	m3 := map[string]map[string]string{
		"basil": map[string]string{
			"fistName": "Vasiliy",
			"lastName": "Zhurin",
		},
	}
	fmt.Println(m3)

	if value, ok := m3["basil"]; ok {
		fmt.Println(value)
	}
}

func findMin(list []int) int {
	if len(list) == 0 {
		return 0
	}

	min := list[0]

	for _, value := range list {
		if value < min {
			min = value
		}
	}

	return min
}

func function() {
	m1 := multiply1(2, 2)
	fmt.Println(m1)

	m2 := multiply2(2, 2)
	fmt.Println(m2)

	m3, m4 := multiMultiply1(2, 5, 10)
	fmt.Println(m3, m4)

	m5, m6 := multiMultiply2(2, 5, 10)
	fmt.Println(m5, m6)
}

func multiply1(value int, multiplier int) int {
	return value * multiplier
}

func multiply2(value int, multiplier int) (result int) {
	result = value * multiplier

	return
}

func multiMultiply1(value int, multiplier1 int, multiplier2 int) (result1 int, result2 int) {
	result1 = value * multiplier1
	result2 = value * multiplier2

	return
}

func multiMultiply2(value int, multiplier1 int, multiplier2 int) (int, int) {
	return value * multiplier1, value * multiplier2
}

func variadicFunction() {
	terms := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(terms...))
}

func sum(terms ...int) int {
	sum := 0

	for _, term := range terms {
		sum += term
	}

	return sum
}

func closure() {
	generator := createEvenGenerator()
	fmt.Println(generator())
	fmt.Println(generator())
}

func createEvenGenerator() func() uint {
	x := uint(0)

	return func() (even uint) {
		even = x
		x += 2
		return
	}
}

func recursion() {
	fmt.Println(factorial(5))
}

func factorial(value uint) uint {
	if value == 1 {
		return 1
	}

	return value * factorial(value-1)
}

func defer_() {
	file := openFile()
	defer closeFile(file)
}

func openFile() int {
	id := 1
	fmt.Printf("File with id '%v' opened\n", id)

	return id
}

func closeFile(id int) {
	fmt.Printf("File with id '%v' closed\n", id)
}

func panicAndRecover() {
	defer func() {
		message := recover()
		fmt.Println(message)
	}()
	panic("PANIC")
}

func struct_() {
	var c1 Circle1
	fmt.Println(c1)

	var pc2 = new(Circle1)
	fmt.Println(pc2)

	pc3 := new(Circle1)
	fmt.Println(pc3)

	c4 := Circle1{0, 0, 7}
	fmt.Println(c4)
	fmt.Println("Area:", c4.area())

	c5 := Circle1{x: 0, y: 0, r: 5}
	fmt.Println(c5)

	pc6 := &Circle1{0, 0, 5}
	fmt.Println(pc6)
}

type Circle1 struct {
	x int
	y int
	r int
}

func (c *Circle1) area() float64 {
	return math.Pi * float64(c.r*c.r)
}

func (c *Circle1) perimeter() float64 {
	return 2 * math.Pi * float64(c.r)
}

type Circle2 struct {
	x, y, d int
}

func (c *Circle2) perimeter() float64 {
	return math.Pi * float64(c.d)
}

func (c *Circle2) area() float64 {
	return math.Pi * float64(c.d*c.d) / 4
}

func embeddedType() {
	type Person struct {
		name string
	}

	type Android struct {
		Person
		serialNumber string
	}

	android := Android{Person{name: "A"}, "7"}
	fmt.Println(android.Person.name)
	fmt.Println(android.name)
	fmt.Println(android.serialNumber)
}

func interface_() {
	c1 := Circle1{0, 0, 5}
	c2 := Circle2{5, 5, 10}
	fmt.Println(totalShapesArea([]Shape{&c1, &c2}))
	fmt.Println(totalShapesPerimeter([]Shape{&c1, &c2}))
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalShapesArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}

	return total
}

func totalShapesPerimeter(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.perimeter()
	}

	return total
}

func main() {
	variable()
	loop()
	switch_()
	array()
	slice()
	map_()
	fmt.Println(findMin([]int{4, 3, 1, 2}))
	function()
	variadicFunction()
	closure()
	recursion()
	defer_()
	panicAndRecover()
	struct_()
	embeddedType()
	interface_()
}
