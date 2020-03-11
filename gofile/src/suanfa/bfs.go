package main
import (
	"fmt"
)

func bfs1(grid [][]int) int{
	row := len(grid)
	col := len(grid[0])
	var queue []int
	fresh := 0
	time := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, i, j)
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	
	for {
		count := 0
		size := len(queue)
		if size == 0 {
			break
		}
		for i := 0; i < size; i+=2 {
			m, n := queue[i], queue[i+1]
			if m < row-1 && grid[m+1][n] == 1 {
				grid[m+1][n] = 2
				queue = append(queue, m+1, n)
				fresh--
				count++
			}
			if m > 0 && grid[m-1][n] == 1 {
				grid[m-1][n] = 2
				queue = append(queue, m-1, n)
				fresh--
				count++
			}
			if n < col-1 && grid[m][n+1] == 1 {
				grid[m][n+1] = 2
				queue = append(queue, m, n+1)
				fresh--
				count++
			}
			if n > 0 && grid[m][n-1] == 1 {
				grid[m][n-1] = 2
				queue = append(queue, m, n-1)
				fresh--
				count++
			}
		}
		queue = queue[size:]
		if count > 0 {
			time++
		}
	}
	if fresh > 0 {
		return -1
	}
	return time
}

func bfs2(grid [][]int, v int) {
	queue := []int{}
	queue = append(queue, 0)
	visit := map[int]bool{}
	l := len(grid)
	for len(queue) > 0 {
		q := queue[0]
		fmt.Println(q)
		queue = queue[1:]
		for i := 0; i < l; i++ {
			if grid[q][i] == 1 {
				if !visit[i] {
					queue = append(queue, i)
				}
				visit[i] = true
			}
		}
	}
}

func dfs(grid [][]int, v int) {
	
}

func initGraph1(grid [][]int) [][]int {
	grid = [][]int{{2,1,1},{1,1,0},{0,1,1}}
	return grid
}

func initGraph2(grid [][]int) [][]int {
	grid = [][]int{{0,0,1,0},{0,0,0,0},{0,0,0,1},{0,1,1,0}}
	return grid
}

// func initGraph2(grid []map) {
// 	graph := make(map[byte] []byte)
// 	graph['A'] = []byte{'B', 'C'}
// 	graph['B'] = []byte{'A', 'C', 'D'}
// 	graph['C'] = []byte{'A', 'B', 'D', 'E'}
// 	graph['D'] = []byte{'B', 'C', 'E', 'F'}
// 	graph['E'] = []byte{'C', 'D'}
// 	graph['F'] = []byte{'D'}
// }

func main() {
	grid := [][]int{}
	grid = initGraph1(grid)
	fmt.Println(bfs1(grid))
	grid = initGraph2(grid)
	bfs2(grid, 0)
}
