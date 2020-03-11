package main
import (
    "fmt"
)

func orangesRotting(grid [][]int) int {
    type zb struct {x, y int}
    f := 0
    minute := 0
    m := make(map[zb] bool)

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                f ++
                m[zb{i, j}] = true
            }
        }
    }

    for f != 0 {
        f = 0
        fl := []zb{}
        for k, _ := range m {
            fl = append(fl, k)
        }
        for i := 0; i < len(fl); i++ {
            if fl[i].x + 1 < len(grid) && grid[fl[i].x+1][fl[i].y] == 1 {
                grid[fl[i].x+1][fl[i].y] = 2
                m[zb{fl[i].x+1, fl[i].y}] = true
                f++
            }
            if fl[i].x - 1 >= 0 && grid[fl[i].x-1][fl[i].y] == 1 {
                grid[fl[i].x-1][fl[i].y] = 2
                m[zb{fl[i].x-1, fl[i].y}] = true
                f++
            }
            if fl[i].y + 1 < len(grid[0]) && grid[fl[i].x][fl[i].y+1] == 1 {
                grid[fl[i].x][fl[i].y+1] = 2
                m[zb{fl[i].x, fl[i].y+1}] = true
                f++
            }
            if fl[i].y - 1 >= 0 && grid[fl[i].x][fl[i].y-1] == 1 {
                grid[fl[i].x][fl[i].y-1] = 2
                m[zb{fl[i].x, fl[i].y-1}] = true
                f++
            }
        }
        if f != 0 {
            minute ++
        }
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid); j++ {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }
    return minute
}

func main() {
    grid := [][]int{{2,1,1},{1,1,0},{0,1,1}}
    fmt.Println(orangesRotting(grid))   
}

