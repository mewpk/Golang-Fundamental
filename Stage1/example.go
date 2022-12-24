package main

import (
	"fmt"
)

// ------------ Struct-----------

type Person struct {
	name   string
	weight int
	height float32
}

func main() {
	// --------------- Variable ----------------
	var name string
	name = "Mew PK" // Manual Type Declaration
	age := 200      // Type Inference
	var score float32 = 21.321
	const codeStatus string = "404"
	fmt.Println(name, age, score, codeStatus)
	fmt.Printf("My name is %v \n", name) // %v = value
	fmt.Printf("Score is %T \n", score)  // %T = type
	// ------------- Input ---------------
	var point int
	fmt.Print("type ur point:")
	fmt.Scanf("%d", &point)
	fmt.Println("Point = ", point)

	// --------------- if condition ----------
	if point > 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	//----------------- array ---------------
	var setNumber [4]int = [4]int{1, 2, 3, 4}
	setText := []string{"Cat", "Dog"}

	fmt.Println(setNumber, len(setNumber))
	fmt.Println(setText, len(setText))
	//---------------Slices------------
	fmt.Println(setNumber[:2])
	// --------------Map---------------
	country := map[string]string{}
	country["TH"] = "Thailand"

	value, check := country["JP"]

	if check {
		fmt.Println(value)
	} else {
		fmt.Println("not found!")
	}
	//-----------loop --------------
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//------------ function-----------
	fmt.Println(sum(10, 20))
	fmt.Println(sumAll(10, 20, 30, 40, 50, 60))
	// ------------ Struct-----------
	Mewpk := Person{name: "MewPk", weight: 50, height: 160.5}
	fmt.Println(Mewpk)
}

func sum(a, b int) int {
	return a + b
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
