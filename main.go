package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/astar"
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
	startState := search.NewState(board, goal)
	solution, frontier, expanded := astar.Solve(startState, goal)
	if solution == nil {
		fmt.Println("Çözülemedi...")
		fmt.Println("Frontier'e Giren Düğüm Sayısı:", frontier)
		fmt.Println("Frontier'den Çıkan Düğüm Sayısı:", expanded)
		fmt.Println("End...", time.Now())
		return
	}
	fmt.Println("Çözüm Maliyeti:", solution.NumMoves)
	fmt.Println("Frontier'e Giren Düğüm Sayısı:", frontier)
	fmt.Println("Frontier'den Çıkan Düğüm Sayısı:", expanded)
	fmt.Println("End...", time.Now())
	return
}
