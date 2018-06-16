package dfs

import (
	"testing"
)

func TestDfs(t *testing.T) {
	tt := &[]struct {
		name string
		grid map[int][]int
		res  int
	}{
		{"rigid", map[int][]int{
			0: []int{0, 0, 2, 2},
			1: []int{0, 1, 1, 2},
			2: []int{1, 1, 2, 2}}, 5,
		},
		{"flexible", map[int][]int{
			0: []int{0, 0, 1, 1},
			1: []int{0, 1, 2, 1, 2},
			2: []int{1, 1, 2, 2, 2, 1},
			3: []int{0, 0, 2, 1, 2, 0, 3}}, 7,
		},
		{"large", map[int][]int{
			0:  []int{0, 0, 1, 1, 2, 3, 3, 1, 2, 1, 1},
			1:  []int{0, 0, 1, 1, 2, 1, 3, 2, 2, 2, 1},
			2:  []int{2, 0, 1, 1, 1, 3, 1, 1, 2, 1, 1},
			3:  []int{0, 2, 1, 1, 2, 3, 3, 2, 2, 1, 2},
			4:  []int{0, 0, 2, 1, 1, 1, 1, 1, 1, 1, 1},
			5:  []int{1, 1, 1, 1, 2, 3, 3, 2, 2, 2, 1},
			6:  []int{0, 0, 1, 2, 2, 3, 3, 2, 2, 1, 1},
			7:  []int{0, 0, 1, 1, 2, 3, 3, 1, 2, 1, 1},
			8:  []int{3, 0, 2, 1, 1, 3, 3, 2, 1, 2, 1},
			9:  []int{0, 3, 1, 1, 2, 1, 3, 2, 2, 1, 1},
			10: []int{3, 0, 1, 1, 2, 3, 1, 1, 2, 2, 1},
			11: []int{0, 0, 1, 1, 2, 3, 3, 2, 1, 1, 2}}, 47,
		},
	}

	for _, tc := range *tt {
		t.Run(tc.name, func(t *testing.T) {
			if r := Do(tc.grid); r != tc.res {
				t.Fatalf("expected %v; got %v", tc.res, r)
			}
		})
	}
}
