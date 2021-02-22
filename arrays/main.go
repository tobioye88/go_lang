package main

import "fmt"

func main() {
	// initialize array
	var fruitArray [2]string
	
	//assign values
	fruitArray[0] = "Apples"
	fruitArray[1] = "Oranges"

	// declare and assign
	newFruitArray := [2]string{"Apples", "Oranges"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])
	fmt.Println(newFruitArray)
	
	// Slices
	sliceFruitArray := []string{"Apples", "Oranges", "Bananas"}
	fmt.Println(sliceFruitArray)

	//length 
	fmt.Println(len(sliceFruitArray))
	fmt.Println(sliceFruitArray[1:2])
}