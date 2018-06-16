package dfs_test

import (
	"fmt"

	"github.com/julesmike/examples/dfsgrid/dfs"
)

// Demonstrating DFS with a 3x4 grid of integers
// ranging from 0 to 2
func ExampleDo() {
	grid := map[int][]int{
		0: []int{0, 0, 2, 2},
		1: []int{0, 1, 1, 2},
		2: []int{1, 1, 2, 2},
	}

	max := dfs.Do(grid)
	fmt.Println("Max connected integer is", max)
	// Output:
	// Max connected integer is 5
}
