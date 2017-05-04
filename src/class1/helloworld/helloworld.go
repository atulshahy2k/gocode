package main

import (
	"fmt"
	"class1/counter"
)


func main() {

	str := "Atul"
	if  counter.CountEven(str) {
		fmt.Printf("%s -  Even Fellow.\n",str)

	} else {
		fmt.Printf("%s -  Odd Fellow .\n",str)
	}

}


