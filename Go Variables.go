package main

import (
	"fmt"
)

func main() {
	var student1 string = "Nilay" //type is string
	var student2 = "Rakib"        //type is inferred
	x := 2                        //type is int
	y := 3.5                      //type is float
	t := x + int(y)               //here y is int so y=3. you need to use same type of values for +,-,* or /
	k := float64(x) + y           //here x is float and x= 2.0
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(t) //here t is int so the total value is 5(int)
	fmt.Println(k) //here k is float so the total value is 5.5(float)
}
