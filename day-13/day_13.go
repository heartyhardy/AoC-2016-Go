package main

const secretKey = 1362

const width = 50
const height = 50

const (
	blocked = iota
	pathable
)

//Cell ...
type Cell struct {
	x, y, flag uint8
}

func scanGrid() {

}

func isPathable(x, y uint16) bool {
	var on, off uint8
	var sum uint16 = (x*x + 3*x + 2*x*y + y + y*y) + secretKey
	for i := 0; i < 15; i++ {
		if bit := sum & (1 << i); bit != 0 {
			on++
		}
	}
	return on
}

func main() {

}
