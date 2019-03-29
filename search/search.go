package search

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/utils"
	"math"
)

////////STACK DEFs
type Stack []([][]int)

func (s Stack) Empty() bool { return len(s) == 0 }

//func (s Stack) Peek() int      { return s[len(s)-1] }
func (s *Stack) Put(i [][]int) { (*s) = append((*s), i) }
func (s *Stack) Pop() [][]int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

/////////END STACK DEFs
type State struct {
	Board    [][]int
	NumMoves int
	Parent   *State
	LastMove Direction
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

//////COMPARES

func FindEmptyTile(board [][]int) (int, int) {
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			if board[y0][x0] == 0 {
				return x0, y0
			}
		}
	}
	return -1, -1
}

////////ACTIONS
func MoveUp(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyY > 0 {
		newBoard[emptyY-1][emptyX], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY-1][emptyX]
		return newBoard, emptyX, emptyY - 1
	}
	return board, emptyX, emptyY
}

func MoveDown(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyY < 2 {
		newBoard[emptyY+1][emptyX], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY+1][emptyX]
		return newBoard, emptyX, emptyY + 1
	}
	return board, emptyX, emptyY

}

func MoveLeft(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyX > 0 {
		newBoard[emptyY][emptyX-1], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY][emptyX-1]
		return newBoard, emptyX - 1, emptyY
	}
	return board, emptyX, emptyY

}

func MoveRight(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyX < 2 {
		newBoard[emptyY][emptyX+1], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY][emptyX+1]
		return newBoard, emptyX + 1, emptyY
	}
	return board, emptyX, emptyY

}

//check possible moves to create children
func (s State) PossibleMoves() []State {
	x, y := FindEmptyTile(s.Board)
	fmt.Println(x, y)
	moves := make([]State, 0, 4)
	fmt.Println("possiblemovesstart")
	if y > 0 {
		fmt.Println("canmoveup")

		newBoard3, _, _ := MoveUp(s.Board, x, y)
		moves = append(moves, State{Board: newBoard3, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Up})
		utils.StatePrinter(newBoard3)
	}
	if y < 2 {
		fmt.Println("canmovedown")

		newBoard4, _, _ := MoveDown(s.Board, x, y)
		moves = append(moves, State{Board: newBoard4, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Down})
		utils.StatePrinter(newBoard4)
	}
	if x > 0 {
		fmt.Println("canmoveleft")

		newBoard1, _, _ := MoveLeft(s.Board, x, y)
		moves = append(moves, State{Board: newBoard1, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Left})
		utils.StatePrinter(newBoard1)
	}
	if x < 2 {
		fmt.Println("canmoveright")
		newBoard2, _, _ := MoveRight(s.Board, x, y)
		moves = append(moves, State{Board: newBoard2, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Right})
		utils.StatePrinter(newBoard2)
	}
	utils.StatePrinter(s.Board)
	fmt.Println("possiblemovesend")
	return moves

}

////BEGINING & END
func NewState(board [][]int) State {
	return State{Board: board}
}

func (s State) IsGoal(goal [][]int) bool {
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			if s.Board[y0][x0] != goal[y0][x0] {
				return false
			}
		}
	}
	return true
}

/////HEURISTIC METHODS
func ManhattanDistance(board [][]int, goal [][]int) int {
	var dx, dy int
	var sum int
	sum = 0
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			for y1 := 0; y1 < 3; y1++ {
				for x1 := 0; x1 < 3; x1++ {
					if board[y0][x0] == goal[y1][x1] && board[y0][x0] != 0 {
						dx = int(math.Abs(float64(x0 - x1)))
						dy = int(math.Abs(float64(y0 - y1)))
						sum = sum + dx + dy
						//fmt.Printf("x0: %d, y0:%d, sta: %d,,y0:%d,x0:%d,goa:%d\n", x0, y0, board[y0][x0], x1, y1, goal[y1][x1])
						//fmt.Printf("dx:%d,dy:%d,sum:%d\n", dx, dy, sum)
					}
				}
			}

		}
	}
	return sum
}
