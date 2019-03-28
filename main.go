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
	fmt.Println("End...", time.Now())
}
