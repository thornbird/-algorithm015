func minPathSum(grid [][]int) int {
	height := len(grid)

	if height == 0 {
		return 0
	}

	width := len(grid[0])
	if width == 0 {
		return 0
	}

	for i := 1; i < height; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < width; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[height-1][width-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}