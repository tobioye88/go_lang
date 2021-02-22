package main

import "fmt"


func main(){
	// string
	// bool
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32, float 64
	// complex64, complex128

	var i int = 2
	var name string = "John doe"

	const isCool = true

	// short hand
	job := "Software Engineer"
	// multiple creation
	age, height, hobby := 20, 1.5, "swiming"

	fmt.Println(i, name, job, age, height, hobby)
}