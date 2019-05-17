package main

import (
	"fmt"
	"math"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func deferme() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

// vardiadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Closures and anonymous functions
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Maps
type Vertex struct {
	X, Y float64
}

// Methods
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// Maps Wordcount function
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	for i := range words {
		wordCount[words[i]]++
	}
	return wordCount
}

func main() {
	// fmt.Println("hey!")
	// var b, c int = 1, 2
	// fmt.Println(b, c)
	// fmt.Println(add(b, c))

	// response, err := http.Get("http://google.com/")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// 	return
	// }
	// fmt.Println(response)

	// y, z := swap("hello", "world")
	// fmt.Println(y, z)

	// Defer
	// defer fmt.Println("world")
	// fmt.Println("hello")
	// deferme()
	// fmt.Println("counting")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	// fmt.Println("done")

	// Slices
	// A slice does not store any data, it just describes a section of an underlying array
	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	// s := make([]string, 3)
	// fmt.Println("empty:", s)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[1])
	// fmt.Println(len(s))
	// s = append(s, "d")
	// fmt.Println(s)
	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("copy:", c)

	// Square root
	// fmt.Println(Sqrt(2))
	// fmt.Println("break")
	// fmt.Println(math.Sqrt(2))

	// Range provides both the index and value for each entry
	// The index value can be ignored with the _ identifier
	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum:", sum)
	// pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	// for i, v := range pow {
	// 	fmt.Printf("2**%d=%d\n", i, v)
	// }

	// Variadic functions
	// sum(1, 2)
	// sum(1, 2, 3)
	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	// Closures and anonymous functions
	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// newInt := intSeq()
	// fmt.Println(newInt())

	// Recursion
	// fmt.Println(factorial(5))

	// WordCount
	// s := "I am am learning Go!"
	// fmt.Println(WordCount(s))

	// Methods
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())

}

