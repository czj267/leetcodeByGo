package main

import (
	"fmt"
	"os"
)

/**
读取迷宫数据
*/
func mazeData(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, err = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type Point struct {
	x, y int
}

/**
判断点是否在迷宫内
*/
func (p Point) at(maze [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(maze) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(maze[0]) {
		return 0, false
	}
	return maze[p.x][p.y], true
}

/**
移动点的方向
*/
func (p Point) add(r Point) Point {
	return Point{x: p.x + r.x, y: p.y + r.y}
}

/**
定义四个方向
*/
var dirs = [4]Point{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

func walkMaze(maze [][]int, start, end Point) [][]int {
	steps := make([][]int, len(maze))
	for i, j := range steps {
		fmt.Println(j)
		steps[i] = make([]int, len(maze[0]))
	}

	q := []Point{start}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			p := cur.add(dir)

			val, ok := p.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = p.at(steps)
			if !ok || val != 0 {
				continue
			}
			if p == start {
				continue
			}
			curVal, _ := cur.at(steps)
			steps[p.x][p.y] = curVal + 1
			q = append(q, p)

		}

	}

	return steps
}

func main() {
	steps1 := make([][]int, 5)
	//steps1[1][1] = 6
	fmt.Println(steps1)
	//steps1[0]
	for i, j := range steps1 {
		fmt.Println(i, j)
		steps1[i] = make([]int, 6)
	}

	os.Exit(0)
	maze := mazeData("myTest/maze/mazeData.txt")
	//fmt.Printf("%v", maze)
	steps := walkMaze(maze, Point{0, 0}, Point{6, 5})
	for _, v := range steps {
		for _, p := range v {
			fmt.Printf("%v   ", p)
		}
		fmt.Println()
	}
	//fmt.Println(steps)
}
