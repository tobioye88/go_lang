package main

import "fmt"

func main() {
	//making a map
	maps := make(map[string]string)

	maps["Bob"] = "bob@gmail.com"
	maps["Sharon"] = "sharon@gmail.com"
	maps["Dave"] = "dave@gmail.com"

	fmt.Println(maps)
	fmt.Println(len(maps))
	fmt.Println(maps["Dave"])
	
	//delete from map
	delete(maps, "Sharon")
	fmt.Println(maps)
	
	
	// Declear and assign
	users := map[string]string {"Tobi": "tobi@gmail.com"}
	fmt.Println(users)
}