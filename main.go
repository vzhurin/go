package main

import "fmt"

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

func switchCase() {
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

func main() {
	variable()
	loop()
	switchCase()
	array()
	slice()
}
