package main

import "fmt"

/**
* @name 
*/
type Person struct {
	name string
	gender string
	age int

	// short hand
	// name, gender string
	// age int
}

// method value receiver 
func (person Person) greet() string {
	return "Hi, i am " + person.name + " and i am " + fmt.Sprintf("%d", person.age) + " years old"
}

// method pointer receiver 
func (person *Person) celebrateBirthday(){
	person.age++;
}

func main() {
	john := Person{ name: "John Doe", gender: "m", age: 25 }

	fmt.Println(john)
	fmt.Println(john.name)
	fmt.Println(john.greet())

	john.celebrateBirthday()
	fmt.Println(john)
}