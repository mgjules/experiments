// Package dfs implements a simple depth-first search
package dfs

// Do returns the maximum connected similar integer in any given map[int][]int
func Do(grid map[int][]int) int {
	visited := make([][]bool, len(grid))
	for i, row := range grid {
		visited[i] = make([]bool, len(row))
	}

	max := 0

	for i, row := range grid {
		for j, cell := range row {
			if m := rec(grid, cell, i, j, new(int), visited); m > max {
				max = m
			}
		}
	}

	return max
}

func rec(g map[int][]int, c, i, j int, h *int, v [][]bool) int {
	// check if 1) out of bounds 2) same as previous cell 3) already visited
	if i < 0 || j < 0 || i >= len(v) || j >= len(v[i]) || c != g[i][j] || v[i][j] {
		return *h
	}

	v[i][j] = true
	*h++

	// can be adapted to check all 8 sides instead of 4
	dd := map[string]struct {
		i, j int
	}{
		"right":  {i: 0, j: +1},
		"bottom": {i: +1, j: 0},
		"left":   {i: 0, j: -1},
		"top":    {i: -1, j: 0},
	}

	for _, d := range dd {
		rec(g, c, i+d.i, j+d.j, h, v)
	}

	return *h
}
