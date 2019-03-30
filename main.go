package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/astar"
	"github.com/ozgurOdun/vbm688_hw01/search"
	"github.com/ozgurOdun/vbm688_hw01/utils"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: hw01 bfs or hw01 astar!")
		return
	}
	if strings.Compare(args[1], "bfs") == 0 {
		fmt.Println("bfs")
	} else if strings.Compare(args[1], "bfs") == 0 {
		fmt.Println("astar")
	} else {
		fmt.Println("Usage: hw01 bfs or hw01 astar!")
		return
	}
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
