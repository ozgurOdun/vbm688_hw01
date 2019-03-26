package search

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

	for y0 := 1; y0 < 3; y0++ {
		for x0 := 1; x0 < 3; x0++ {

			for y1 := 1; y1 < 3; y1++ {
				for x1 := 1; x1 < 3; x1++ {
					if state[y0][x0] == goal[y1][x1] {
						if x0-x1 < 0 {
							dx = x1 - x0
						} else {
							dx = x0 - x1
						}
						if y0-y1 < 0 {
							dy = y1 - y0
						} else {
							dy = y0 - y1
						}
						sum = sum + dx + dy
					}
				}
			}

		}
	}
	return sum
}
