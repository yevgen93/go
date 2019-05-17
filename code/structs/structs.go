package main

import "fmt"

type position struct {
	x, y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {
	p := position{4, 2}
	b := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(b)
}

