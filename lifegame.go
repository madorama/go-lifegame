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

func nextLife(cell CellStatus, lives uint) CellStatus {
	switch cell {
	case Dead:
		if lives == 3 {
			return Alive
		}
	case Alive:
		if lives <= 1 || lives >= 4 {
			return Dead
		}
	}
	return cell
}

func (game *Lifegame) rule(x uint, y uint) CellStatus {
	var lives uint = 0
	pos := []int{-1, 0, 1}
	for _, iy := range pos {
		for _, ix := range pos {
			if ix == 0 && iy == 0 {
				continue
			}
			px, py := int(x)+ix, int(y)+iy
			if px < 0 || px >= int(game.width) || py < 0 || py >= int(game.height) {
				continue
			}
			lives += uint(game.world[py][px])
		}
	}
	cell := game.world[y][x]
	return nextLife(cell, lives)
}

func (game *Lifegame) update() {
	newWorld := make([][]CellStatus, game.height)
	for i := range newWorld {
		newWorld[i] = make([]CellStatus, game.width)
	}

	for iy := range newWorld {
		for ix := range newWorld[iy] {
			newWorld[iy][ix] = game.rule(uint(ix), uint(iy))
		}
	}
	game.world = newWorld
}

func main() {
	var game = *randomInit(50, 50)
	fmt.Println(&game.world)
}
