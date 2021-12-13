package main

import "fmt"

func main() {
	var a string
	var b int
	var c bool
	//These variables are declared but they have not been assigned initial values.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//By running the code, we can see that they already have the default values of their respective types
}
