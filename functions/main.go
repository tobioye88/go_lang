package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func add(first, second int) int {
	return first + second
}

func main(){
	fmt.Println(greeting("Tobi"))
	fmt.Println(add(1, 2))
}