package search

import (
	_ "fmt"
	"math"
)

type State struct {
	Board     [][]int
	EmptyTile int
	NumMoves  int
	Children  []*State
	Parent    *State
	LastMove  Direction
	Distance  int
}

type Direction int

const (
	None  Direction = 0
	Up    Direction = 1
	Down  Direction = 2
	Left  Direction = 3
	Right Direction = 4
)

const SizeX = 3
const SizeY = 3

func ManhattanDistance(state [][]int, goal [][]int) int {
	var dx, dy int
	var sum int
	sum = 0
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			for y1 := 0; y1 < 3; y1++ {
				for x1 := 0; x1 < 3; x1++ {
					if state[y0][x0] == goal[y1][x1] && state[y0][x0] != 0 {
						dx = int(math.Abs(float64(x0 - x1)))
						dy = int(math.Abs(float64(y0 - y1)))
						sum = sum + dx + dy
						//fmt.Printf("x0: %d, y0:%d, sta: %d,,y0:%d,x0:%d,goa:%d\n", x0, y0, state[y0][x0], x1, y1, goal[y1][x1])
						//fmt.Printf("dx:%d,dy:%d,sum:%d\n", dx, dy, sum)
					}
				}
			}

		}
	}
	return sum
}
