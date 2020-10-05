func islandPerimeter(grid [][]int) int {
    perimeter := 0
    height := len(grid)
    width := len(grid[0])

    if height == 0 || width == 0 {
        return perimeter
    }

    for i, column := range grid {
        for j, item := range column {
            if item == 0 {
                continue
            }
            perimeter += 4
            if i >= 1 && grid[i-1][j] == 1 {
                perimeter--
            }
            if i <= height-2 && grid[i+1][j] == 1 {
                perimeter--
            }
            if j >= 1 && grid[i][j-1] == 1 {
                perimeter--
            }
            if j <= width - 2 && grid[i][j+1] == 1{
                perimeter--
            }
        }
    }
    return perimeter
}