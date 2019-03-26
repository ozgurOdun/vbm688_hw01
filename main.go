package main

import (
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/search"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func fillGoal() [][]int {
	var g [][]int
	row1 := []int{0, 1, 2}
	row2 := []int{3, 4, 5}
	row3 := []int{6, 7, 8}
	g = append(g, row1)
	g = append(g, row2)
	g = append(g, row3)
	return g
}
func main() {
	fmt.Println("Starting...", time.Now())
	var board [][]int
	goal := make([][]int, 9)
	goal = fillGoal()
	file, err := os.Open("input.txt")
	checkErr(err)

	rawInput := make([]byte, 17)
	n, err := file.Read(rawInput)
	if n > 0 {
		fmt.Println("Raw input is: ", string(rawInput))
	}
	var inputArray []int
	var inpStr []string
	inputLines := strings.Split(string(rawInput), "\n")
	fmt.Println(len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		fmt.Printf("Input line %d is %s\n", i, inputLines[i])
		inpStr = strings.Split(inputLines[i], " ")
		fmt.Println("len of inpStr is:", len(inpStr))
		for j := 0; j < len(inpStr); j++ {
			ij, err := strconv.ParseInt(inpStr[j], 0, 32)
			checkErr(err)
			fmt.Println("j is: ", j, inpStr[j], ij)
			inputArray = append(inputArray, int(ij))
			fmt.Println(inputArray)

		}

	}
	row1 := []int{inputArray[0], inputArray[1], inputArray[2]}
	row2 := []int{inputArray[3], inputArray[4], inputArray[5]}
	row3 := []int{inputArray[6], inputArray[7], inputArray[8]}
	board = append(board, row1)
	board = append(board, row2)
	board = append(board, row3)
	fmt.Println(inputArray)
	fmt.Println(board)
	fmt.Println(goal)
	fmt.Println(search.ManhattanDistance(board, goal))
}
