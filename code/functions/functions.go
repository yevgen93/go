package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello", name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABunchOfTimes() {
	fmt.Println("Hello")
	sayHelloABunchOfTimes()
}

func main() {
	hello("Bob")

	x := 5
	x = addOne(x)
	fmt.Println(x)

	x = addOne(addOne(x))
	fmt.Println(x)

	sayHelloABunchOfTimes()
}

