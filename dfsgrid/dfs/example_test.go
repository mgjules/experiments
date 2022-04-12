package dfs_test

import (
	"fmt"

	"github.com/mgjules/toys/dfsgrid/dfs"
)

// Demonstrating DFS with a 3x4 grid of integers
// ranging from 0 to 2
func Example_usage() {
	grid := map[int][]int{
		0: {0, 0, 2, 2},
		1: {0, 1, 1, 2},
		2: {1, 1, 2, 2},
	}

	max := dfs.Do(grid)
	fmt.Println("Max connected integer is", max)
	// Output:
	// Max connected integer is 5
}
