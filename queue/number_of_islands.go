package queue

import "github.com/mngoutham/GOExamples"

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//Example 1:
//
//Input:
//11110
//11010
//11000
//00000
//
//Output: 1

func numIslands(grid [][]byte) int {
	x := len(grid)
	if x == 0 {
		return 0
	}
	y := len(grid[0])
	count := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 0 {
				continue
			}
			markIsland(grid, i, j, func(i, j int) bool {
				if i < 0 || j < 0 || i >= x || j >= y || grid[i][j] == 0 {
					return false
				}
				return true
			})
			count += 1
		}
	}
	return count
}

func markIsland(grid [][]byte, i, j int, isValid func(int, int) bool) {
	grid[i][j] = 0
	for _, incr := range GOExamples.INCR {
		newI := i + incr.i
		newJ := j + incr.j
		if isValid(newI, newJ) {
			markIsland(grid, newI, newJ, isValid)
		}
	}
	return
}
