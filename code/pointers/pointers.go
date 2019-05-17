package main

import "fmt"

func addOneCopy(x int) {
	x = x + 1
}

func addOnePointer(num *int) {
	*num = *num + 1
}

func main() {

	x := 5
	fmt.Println(x)

	var xPointer *int = &x
	fmt.Println(xPointer)

	addOneCopy(x)
	// adds 1 to the copy of x, not the actual value, so it remains 5
	fmt.Println(x)

	addOnePointer(xPointer)
	// adds 1 to the actual value of x by de-referencing the pointer
	fmt.Println(x)
}

