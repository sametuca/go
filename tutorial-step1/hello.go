package main

import (
	"fmt"
)

func main() {
	// message, err := Hello("Gladys")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(message)
	arr := SplitArray("Gladys")
	fmt.Println(arr)
}

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func SplitArray(name string) []string {
	arr := []string{"a", "b", "c"}
	return arr
}
