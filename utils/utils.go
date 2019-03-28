package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func FillGoal() [][]int {
	var g [][]int
	row1 := []int{0, 1, 2}
	row2 := []int{3, 4, 5}
	row3 := []int{6, 7, 8}
	g = append(g, row1)
	g = append(g, row2)
	g = append(g, row3)
	return g
}

func StatePrinter(state [][]int) {
	fmt.Println("_______")
	fmt.Println(state[0])
	fmt.Println(state[1])
	fmt.Println(state[2])
	fmt.Println("-------")
}

func InputParser(fileName string) [][]int {
	var board [][]int
	file, err := os.Open(fileName)
	checkErr(err)

	rawInput := make([]byte, 17)
	n, err := file.Read(rawInput)
	if n > 0 {
		fmt.Println("Raw input is: ", string(rawInput))
	}
	var inputArray []int
	var inpStr []string
	inputLines := strings.Split(string(rawInput), "\n")
	//fmt.Println(len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		//fmt.Printf("Input line %d is %s\n", i, inputLines[i])
		inpStr = strings.Split(inputLines[i], " ")
		//fmt.Println("len of inpStr is:", len(inpStr))
		for j := 0; j < len(inpStr); j++ {
			ij, err := strconv.ParseInt(inpStr[j], 0, 32)
			checkErr(err)
			//fmt.Println("j is: ", j, inpStr[j], ij)
			inputArray = append(inputArray, int(ij))
			//fmt.Println(inputArray)
		}
	}
	row1 := []int{inputArray[0], inputArray[1], inputArray[2]}
	row2 := []int{inputArray[3], inputArray[4], inputArray[5]}
	row3 := []int{inputArray[6], inputArray[7], inputArray[8]}
	board = append(board, row1)
	board = append(board, row2)
	board = append(board, row3)
	return board

}
