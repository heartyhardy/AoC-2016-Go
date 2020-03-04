// package main

// import (
// 	"crypto/md5"
// 	"fmt"
// )

// type Cells [][]*Cell

// type Queue []Cell

// type Coordinate struct {
// 	x int
// 	y int
// 	d string
// }

// type Cell struct {
// 	Coordinate
// 	doors     []bool
// 	hash      string
// 	direction string
// 	locked    bool
// }

// func (q Queue) enqueue(c *Cell) Queue {
// 	q = append(q, *c)
// 	return q
// }

// func (q Queue) peek() *Cell {
// 	if len(q) > 0 {
// 		return &q[0]
// 	}
// 	return nil
// }

// func (q Queue) dequeue() Queue {
// 	t := make(Queue, len(q)-1)
// 	copy(t, q[1:])
// 	q = nil
// 	q = t
// 	return q
// }

// func buildGrid(size int, hash string) Cells {
// 	grid := make(Cells, size)
// 	for r := 0; r < size; r++ {
// 		grid[r] = make([]*Cell, size)
// 		for c := 0; c < size; c++ {
// 			ncell := new(Cell)
// 			ncell.x, ncell.y = c, r
// 			ncell.doors = make([]bool, 4)
// 			grid[r][c] = ncell
// 		}
// 	}
// 	grid[0][0].hash = hash
// 	return grid
// }

// func getShortestPath(grid Cells) {
// 	start, end := grid[0][0], grid[3][3]
// 	queue := make(Queue, 0)
// 	queue = queue.enqueue(start)

// 	for len(queue) > 0 {
// 		next := queue.peek()
// 		queue = queue.dequeue()

// 		if next.x == end.x && next.y == end.y {
// 			fmt.Println(next.x, next.y, next.hash, next.direction)
// 			break
// 		}

// 		queue = scanNeighbors(grid, queue, next)
// 	}
// }

// func scanNeighbors(grid Cells, q Queue, c *Cell) Queue {
// 	adjacent := map[int]Coordinate{
// 		0: Coordinate{x: 0, y: -1, d: "U"},
// 		1: Coordinate{x: 0, y: 1, d: "D"},
// 		2: Coordinate{x: -1, y: 0, d: "L"},
// 		3: Coordinate{x: 1, y: 0, d: "R"},
// 	}
// 	hash := fmt.Sprintf("%x", md5.Sum([]byte(c.hash+c.direction)))[:4]
// 	for k, h := range hash {
// 		if h > 97 && h < 103 {
// 			dx := c.x + adjacent[k].x
// 			dy := c.y + adjacent[k].y
// 			if dx >= 0 && dy >= 0 && dx < 4 && dy < 4 {

// 				if !areAllLocked(c.hash + c.direction + adjacent[k].d) {
// 					grid[dy][dx].hash = c.hash
// 					grid[dy][dx].direction = c.direction + adjacent[k].d
// 					q = q.enqueue(grid[dy][dx])
// 				}
// 				if dy == 3 && dx == 3 {
// 					grid[dy][dx].hash = c.hash
// 					grid[dy][dx].direction = c.direction + adjacent[k].d
// 					q = q.enqueue(grid[dy][dx])
// 				}

// 			}
// 		}
// 	}
// 	return q
// }

// func areAllLocked(hash string) bool {
// 	keys := map[rune]bool{
// 		'b': true,
// 		'c': true,
// 		'd': true,
// 		'e': true,
// 		'f': true,
// 	}
// 	md5hash := fmt.Sprintf("%x", md5.Sum([]byte(hash)))[:4]
// 	for _, r := range md5hash {
// 		if _, ok := keys[r]; ok {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	grid := buildGrid(4, "ihgpwlah")
// 	getShortestPath(grid)
// }
