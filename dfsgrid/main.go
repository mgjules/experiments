package main

import (
	"fmt"

	"github.com/mgjules/experiments/dfsgrid/dfs"
)

const (
	red   = iota // 0
	blue         // 1
	green        // 2
)

func main() {
	grid := map[int][]int{
		0: {blue, blue, green, red},
		1: {blue, green, red, green},
		2: {red, green, green, green},
	}

	fmt.Println(dfs.Do(grid))
}
