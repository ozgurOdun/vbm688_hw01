package astar

import (
	"container/heap"
	"fmt"
	"github.com/ozgurOdun/vbm688_hw01/priorityQueue"
	"github.com/ozgurOdun/vbm688_hw01/search"
	"github.com/ozgurOdun/vbm688_hw01/utils"
	//"time"
)

func Solve(start search.State, goal [][]int) (*search.State, int, int) {
	var frontier, expanded int
	states := make(map[string]search.State)
	pq := make(priorityQueue.PriorityQueue, 0)
	key := utils.BoardStringer(start.Board)
	states[key] = start
	heap.Push(&pq, &priorityQueue.Item{Value: key, Priority: 0, Index: 0})

	for pq.Len() != 0 {
		//time.Sleep(100 * time.Millisecond)
		currentItem := heap.Pop(&pq).(*priorityQueue.Item)
		current := states[currentItem.Value]
		expanded++
		utils.StatePrinter(current.Board)
		fmt.Println(current.Distance)
		if current.IsGoal(goal) {
			//solved
			return &current, frontier, expanded
		}
		for _, next := range current.PossibleMoves() {
			//implement key of state to keep in heap
			key := utils.BoardStringer(next.Board)
			if old, exists := states[key]; !exists || next.Distance < old.Distance {
				states[key] = next
				heap.Push(&pq, &priorityQueue.Item{Value: key, Priority: next.Distance, Index: 0})
				frontier++
			}
		}
	}
	return nil, frontier, expanded
}
