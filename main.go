package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/search"
	"github.com/ozgurOdun/vbm688_hw01/utils"
	"time"
)

func main() {
	fmt.Println("Starting...", time.Now())
	goal := make([][]int, 9)
	goal = utils.FillGoal()
	board := utils.InputParser("input.txt")
	utils.StatePrinter(board)
	fmt.Println("Manhattan Distance of input is: ", search.ManhattanDistance(board, goal))
	var s1 search.Stack
	s1.Put(board)
	fmt.Println(s1)
	x, y := search.FindEmptyTile(board)
	board1, x, y := search.MoveUp(board, x, y)
	s1.Put(board1)
	fmt.Println(s1[0], s1[1])
	_ = s1.Pop()
	fmt.Println(s1)
	_ = s1.Pop()
	fmt.Println(s1)
	fmt.Println("End...", time.Now())
}
