package main

import (
	"fmt"
)

func main() {
	var student1 string = "Nilay" //type is string
	var student2 = "Rakib"        //type is inferred
	x := 2                        //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
