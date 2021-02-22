package main

import "fmt"


func main() {
	ids := []int {1,2,3,4,5,6}
	for i, v := range ids {
		fmt.Println(i, v)
	}

	myMaps := map[string]string{ "tobi": "tobi@gmail.com", "john": "john@gmail.com", "sam": "sam@gmail.com" }
	for k, v := range myMaps {
		fmt.Printf("%s: %s\n", k, v)
	}

	// for keys
	for k := range myMaps {
		fmt.Printf("Name: %s\n", k)
	}
}