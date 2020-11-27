package main

import "fmt"

func main() {

	var a int // if you are not assigning value, do mention variable type
	//Default value of non-assign int variable is 0..

	var b = 5 // No need to mention variable type if u are assigning value to it

	c := 3 // equals to "var c = 3" , ":" it auto determine variable type

	const d = 6 // you can not change the value of d

	fmt.Println("Hello Vivek !!")
	fmt.Println(a, b, c, d)
}
