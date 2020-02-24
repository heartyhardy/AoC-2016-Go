package main

import (
	"fmt"
)

const secretKey = 1362

const width = 50
const height = 50

const (
	blocked = iota
	pathable
)

//Cell ...
type Cell struct {
	x, y uint8
	flag bool
	BFSdata
}

//BFSdata is Used internally to run BFS
type BFSdata struct {
	path    bool
	parent  *Cell
	visited bool
}

//Queue and its methods
type queue []*Cell

func (q queue) enqueue(c *Cell) queue {
	q = append(q, c)
	return q
}

func (q queue) peek() *Cell {
	return q[0]
}

func (q queue) dequeue() queue {
	t := make([]*Cell, len(q)-1)
	copy(t, q[1:])
	q = nil
	q = t
	return q
}

// End of queue

func scanGrid() {
	var r, c uint16
	grid := make([][]*Cell, height)
	for r = 0; r < height; r++ {
		grid[r] = make([]*Cell, width)
		for c = 0; c < width; c++ {
			newCell := new(Cell)
			newCell.x = uint8(c)
			newCell.y = uint8(r)
			if isPath := isPathable(c, r); isPath {
				newCell.flag = isPath
			}
			grid[r][c] = newCell
		}
	}
	getShortestPath(grid, 31, 39)
	printGrid(grid)
}

func getShortestPath(grid [][]*Cell, tx, ty uint8) {
	locations := 0
	grid[1][1].visited = true
	queue := make(queue, 0)
	queue = queue.enqueue(grid[1][1])

	for len(queue) != 0 {
		next := queue.peek()
		queue = queue.dequeue()
		//fmt.Println(next.x, next.y, queue)
		if next != nil {

			if true {
				grid[next.y][next.x].path = true
				parent := next.parent
				steps := 0

				for parent != nil {
					steps++
					grid[parent.y][parent.x].path = true
					parent = parent.parent
				}
				fmt.Println("Part I: Steps: ", steps)
				if steps == 50 {
					fmt.Println("Locations: ", locations)
					return
				}
			}

			xs := []int{0, 1, 0, -1}
			ys := []int{-1, 0, 1, 0}

			for n := 0; n < 4; n++ {

				dx, dy := int(next.x)+xs[n], int(next.y)+ys[n]

				if (dx >= 0 && dx < width) && (dy >= 0 && dy < height) {
					if grid[dy][dx].visited || !grid[dy][dx].flag {
						continue
					}
					locations++
					grid[dy][dx].visited = true
					grid[dy][dx].parent = grid[next.y][next.x]
					queue = queue.enqueue(grid[dy][dx])
				}
			}
		}
	}
}

func isPathable(x, y uint16) bool {
	var on uint8
	var sum uint16 = (x*x + 3*x + 2*x*y + y + y*y) + secretKey
	for i := 0; i < 15; i++ {
		if bit := sum & (1 << i); bit != 0 {
			on++
		}
	}
	return (on & 1) == 0
}

func printGrid(grid [][]*Cell) {
	var r, c uint16

	for r = 0; r < height; r++ {
		for c = 0; c < width; c++ {
			if grid[r][c].flag {
				if grid[r][c].path {
					fmt.Printf("\u001b[48;5;17m\u001b[38;5;40m%v\u001b[0m\u001b[0m", "◼ ")
					continue
				}
				fmt.Printf("\u001b[48;5;17m\u001b[38;5;18m%v\u001b[0m\u001b[0m", "◼ ")
			} else {
				fmt.Printf("\u001b[48;5;54m\u001b[38;5;53m%v\u001b[0m\u001b[0m", "# ")
			}
		}
		fmt.Println()
	}
}

func main() {
	scanGrid()
}
