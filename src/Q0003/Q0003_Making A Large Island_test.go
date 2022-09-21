package Q0003

import (
	"fmt"
	"testing"
)

type Point struct {
	X int
	Y int
}

var travel = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	pointIsland := make(map[Point]map[Point]struct{}, n * n)
	max := 0

	// get origin island
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			point := Point{
				X: i,
				Y: j,
			}

			curIsland := map[Point]struct{}{
				 point: {},
			}
			pointIsland[point] = curIsland

			if legal(n, i - 1, j) && grid[i - 1][j] == 1 {
				nbIsland := pointIsland[Point{X: i - 1, Y: j}]
				curIsland = mergeIsland(curIsland, nbIsland)
				pointIsland = setIsland(pointIsland, curIsland)
			}

			if legal(n, i, j - 1) && grid[i][j - 1] == 1 {
				nbIsland := pointIsland[Point{X: i, Y: j - 1}]
				curIsland = mergeIsland(curIsland, nbIsland)
				pointIsland = setIsland(pointIsland, curIsland)
			}

			if len(curIsland) > max {
				max = len(curIsland)
			}
		}
	}

	// change 0 to 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			list := make([]Point, 0, 4)
			for _, tr := range travel {
				curI := i + tr[0]
				curJ := j + tr[1]

				if legal(n, curI, curJ) && grid[curI][curJ] == 1 {
					list = append(list, Point{X: curI, Y: curJ})
				}
			}

			// 周边岛屿面积 + 自己
			curArea := compute(pointIsland, list) + 1
			if curArea > max {
				max = curArea
			}
		}
	}

	return max
}

func legal(n, x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}

func mergeIsland(i1, i2 map[Point]struct{}) map[Point]struct{} {
	if len(i1) < len(i2) {
		i1, i2 = i2, i1
	}

	for p := range i2 {
		i1[p] = struct{}{}
	}

	return i1
}

func setIsland(pointIsland map[Point]map[Point]struct{}, island map[Point]struct{}) map[Point]map[Point]struct{} {
	for p := range island {
		pointIsland[p] = island
	}

	return pointIsland
}

// 计算岛屿面积和
func compute(pointIsland map[Point]map[Point]struct{}, list []Point) int {
	islands := make([]map[Point]struct{}, 0, len(list))

	for _, p := range list {
		island := pointIsland[p]

		ok := true
		for _, is := range islands {
			if isSame(island, is) {
				ok = false
				break
			}
		}

		if ok {
			islands = append(islands, island)
		}
	}

	res := 0
	for _, island := range islands {
		res += len(island)
	}

	return res
}

func isSame(i1, i2 map[Point]struct{}) bool {
	p := Point{X: -1}
	i1[p] = struct{}{}
	_, res := i2[p]
	delete(i1, p)
	return res
}

func Test(t *testing.T) {
	grid := [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(largestIsland(grid))
}

