package main

import (
	"fmt"

	"github.com/julesmike/toys/dfsgrid/dfs"
)

const (
	red   = iota // 0
	blue         // 1
	green        // 2
)

func main() {
	grid := map[int][]int{
		0: []int{blue, blue, green, red},
		1: []int{blue, green, red, green},
		2: []int{red, green, green, green},
	}

	fmt.Println(dfs.Do(grid))
}
