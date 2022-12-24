package main

import "fmt"

func main() {
	// ---------- variable ------------------
	var name string
	name = "Mew PK" // Manual Type Declaration
	age := 200      // Type Inference
	var score float32 = 21.321
	const codeStatus string = "404"
	fmt.Println(name, age, score, codeStatus)
	fmt.Printf("My name is %v \n", name) // %v = value
	fmt.Printf("Score is %T", score)     // %T = type

}
