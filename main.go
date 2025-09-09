package main

import "fmt"

func main() {
	err := Top3Words("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
