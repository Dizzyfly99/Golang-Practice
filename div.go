package main

import "fmt"

var a, b int

func main() {
	fmt.Printf("Provide 1st value: ")
	fmt.Scanln(&a)
	fmt.Printf("Provide 2nd value: ")
	fmt.Scanln(&b)
	fmt.Println(a / b)
}
