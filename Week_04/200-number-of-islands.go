func numIslands(grid [][]byte) int {
	height := len(grid)

	if height == 0 {
		return 0
	}

	width := len(grid[0])
	if width == 0 {
		return 0
	}

	count := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				effectIslands(i, j, height, width, grid)
				count++
			}
		}
	}

	return count
}

func effectIslands(i, j, height, width int, grid [][]byte) {
	if i >= height || j >= width || i < 0 || j < 0 || grid[i][j] != '1' {
		return
	}
	// 随便mark成一个非'1'的字符即可
	grid[i][j] = '0'
	effectIslands(i-1, j, height, width, grid)
	effectIslands(i+1, j, height, width, grid)
	effectIslands(i, j-1, height, width, grid)
	effectIslands(i, j+1, height, width, grid)
	return
}