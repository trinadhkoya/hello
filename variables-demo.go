package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i = 1
	var j int = 2
	k := 3.

	//value and type of variable
	fmt.Printf("%v and %T \n", i, i)
	fmt.Printf("%v and %T \n", j, j)
	fmt.Printf("%v and %T \n", k, k)

	//string conversion
	var l string
	l = strconv.Itoa(j)
	fmt.Printf(" value of l is %v and type is %T", l, l)

}
