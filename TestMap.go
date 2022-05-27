package main

import "fmt"

func main() {
	MyMap := make(map[string]int)
	MyMap["xmmzzz"] = 12132
	MyMap["abcac"] = 11123

	for username, value := range MyMap {
		fmt.Println(username, value)
	}

	delete(MyMap, "abcac")

	for username, value := range MyMap {
		fmt.Println(username, value)
	}

}
