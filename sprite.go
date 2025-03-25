package main

type Sprite struct {
	X int
	Y int
	C rune
}

func NewSprite(x int, y int, char rune) *Sprite {
	return &Sprite{
		X: x,
		Y: y,
		C: char,
	}
}
