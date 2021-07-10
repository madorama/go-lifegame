package main

import (
	"fmt"
	"math/rand"
)

type CellStatus uint

const Dead CellStatus = 0
const Alive CellStatus = 1

type Lifegame struct {
	world  [][]CellStatus
	width  uint
	height uint
}

func randomInit(width uint, height uint) *Lifegame {
	game := new(Lifegame)
	game.world = make([][]CellStatus, height)

	for iy := 0; iy < int(height); iy++ {
		game.world[iy] = make([]CellStatus, width)
		for ix := 0; ix < int(width); ix++ {
			cell := Alive
			if rand.Float32() >= 0.3 {
				cell = Dead
			}
			game.world[iy][ix] = cell
		}
	}
	game.width = width
	game.height = height
	return game
}

func main() {
	var game = *randomInit(50, 50)
	fmt.Println(&game.world)
}
