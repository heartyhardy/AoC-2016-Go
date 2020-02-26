package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Disc struct {
	index     int
	start     int
	current   int
	positions int
}

func (d *Disc) Tick(t int) int {
	d.current = (d.current + t) % d.positions
	return d.current
}

func (d *Disc) IsOpen() bool {
	return d.current == 0
}

type Simulation struct {
	t        int
	position int
	discs    []*Disc
}

func (s *Simulation) Init(config string) *Simulation {
	s.discs = make([]*Disc, 0)
	setup := func() {
		data, _ := ioutil.ReadFile(config)
		discConfig := strings.Split(string(data), "\n")
		for _, disc := range discConfig {
			fields := strings.Fields(disc)

			ndisc := new(Disc)
			ndisc.index, _ = strconv.Atoi(string(fields[1][1:]))
			ndisc.positions, _ = strconv.Atoi(fields[3])
			ndisc.start, _ = strconv.Atoi(fields[11])
			ndisc.current = ndisc.start

			s.discs = append(s.discs, ndisc)
		}
	}
	setup()
	s.position = 0
	s.t = 0

	return s
}

func (s *Simulation) Run(t int) bool {
	current, position := t, 0

	for position < len(s.discs) {
		current++
		s.discs[position].Tick(current)
		if !s.discs[position].IsOpen() {
			return false
		}
		position++
	}
	return true

}

func (s *Simulation) RunUntil() int {
	t := 0
	var ok = false

	for !ok {
		s.Init("../inputs/day 15.txt")
		ok = s.Run(t)
		t++
	}

	return t - 1
}

func main() {
	//Change Input text to get Part I/ Part II
	s := new(Simulation)
	s.Init("../inputs/day 15.txt")

	t := s.RunUntil()
	fmt.Println("Time (t): ", t)
}
