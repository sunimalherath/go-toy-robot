package robot

import "fmt"

type direction string

const (
	NORTH direction = "north"
	SOUTH direction = "south"
	EAST  direction = "east"
	WEST  direction = "west"
)

type IRobot interface {
	Place(int, int, direction) bool
	Move() bool
	TurnLeft()
	TurnRight()
	PrintReport()
}

type robot struct {
	OnTheTable      bool
	XPosition       int
	YPosition       int
	FacingDirection direction
}

func (r *robot) Place(x, y int, facing direction) bool {
	r.XPosition = x
	r.YPosition = y
	r.FacingDirection = facing
	return true
}

func (r *robot) Move() bool {
	switch r.FacingDirection {
	case NORTH:
		r.YPosition++
		return true
	case SOUTH:
		r.YPosition--
		return true
	case EAST:
		r.XPosition++
		return true
	case WEST:
		r.XPosition--
		return true
	default:
		return false
	}
}

func (r *robot) TurnLeft() {
	switch r.FacingDirection {
	case NORTH:
		r.FacingDirection = WEST
	case SOUTH:
		r.FacingDirection = EAST
	case EAST:
		r.FacingDirection = NORTH
	case WEST:
		r.FacingDirection = SOUTH
	}
}

func (r *robot) TurnRight() {
	switch r.FacingDirection {
	case NORTH:
		r.FacingDirection = EAST
	case SOUTH:
		r.FacingDirection = WEST
	case EAST:
		r.FacingDirection = SOUTH
	case WEST:
		r.FacingDirection = NORTH
	}
}

func (r *robot) PrintReport() {
	fmt.Printf(`robot's X position: %d and Y position: %d, facing: %s`, r.XPosition, r.YPosition, r.FacingDirection)
}

func (r robot) isValidPlacement(x, y int) bool {
	// check if the robot is placed within the 5 x 5 grid
	if x < 0 || x > 4 {
		return false
	}
	if y < 0 || y > 4 {
		return false
	}
	return true
}
