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

func CopySlice(dst [][]int, src [][]int) {
	dst[0][0] = src[0][0]
	dst[0][1] = src[0][1]
	dst[0][2] = src[0][2]
	dst[1][0] = src[1][0]
	dst[1][1] = src[1][1]
	dst[1][2] = src[1][2]
	dst[2][0] = src[2][0]
	dst[2][1] = src[2][1]
	dst[2][2] = src[2][2]

}

func BoardStringer(board [][]int) string {
	str := ""
	for _, i := range board[0] {
		str += strconv.Itoa(i)
	}
	for _, i := range board[1] {
		str += strconv.Itoa(i)
	}
	for _, i := range board[2] {
		str += strconv.Itoa(i)
	}
	return str
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
	board := make([][]int, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
	}
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
			fmt.Println(inputArray)
		}
	}

	board[0] = inputArray[0:3]
	board[1] = inputArray[3:6]
	board[2] = inputArray[6:9]
	return board

}
