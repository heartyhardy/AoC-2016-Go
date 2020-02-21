package main

//Generator represents the RTG
type Generator struct {
	floor  int
	paired bool
	chip   *Chip
}

//Chip represents the Microship
type Chip struct {
	floor  int
	paired bool
	rtg    *Generator
}

//Elevator ...
type Elevator struct {
	floor int
	cargo []*interface{}
}

//State ...
type State struct {
	state []*interface{}
}

//Facility represents Radioisotope Testing Facility
type Facility struct {
	stack      []*State
	generators []*Generator
	chips      []*Chip
	elevator   *Elevator
	steps      int
}
