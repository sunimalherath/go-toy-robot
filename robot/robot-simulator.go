package robot

import (
	"fmt"

	"github.com/sunimalherath/go-toy-robot/util"
)

type IRobot interface {
	Place(int, int, util.Direction) bool
	Move() bool
	TurnLeft()
	TurnRight()
	PrintReport()
}

type robot struct {
	OnTheTable      bool
	XPosition       int
	YPosition       int
	FacingDirection util.Direction
}

func (r *robot) Place(x, y int, facing util.Direction) bool {
	r.XPosition = x
	r.YPosition = y
	r.FacingDirection = facing
	return true
}

func (r *robot) Move() bool {
	switch r.FacingDirection {
	case util.NORTH:
		r.YPosition++
		return true
	case util.SOUTH:
		r.YPosition--
		return true
	case util.EAST:
		r.XPosition++
		return true
	case util.WEST:
		r.XPosition--
		return true
	default:
		return false
	}
}

func (r *robot) TurnLeft() {
	switch r.FacingDirection {
	case util.NORTH:
		r.FacingDirection = util.WEST
	case util.SOUTH:
		r.FacingDirection = util.EAST
	case util.EAST:
		r.FacingDirection = util.NORTH
	case util.WEST:
		r.FacingDirection = util.SOUTH
	}
}

func (r *robot) TurnRight() {
	switch r.FacingDirection {
	case util.NORTH:
		r.FacingDirection = util.EAST
	case util.SOUTH:
		r.FacingDirection = util.WEST
	case util.EAST:
		r.FacingDirection = util.SOUTH
	case util.WEST:
		r.FacingDirection = util.NORTH
	}
}

func (r *robot) PrintReport() {
	fmt.Printf(`robot's X position: %d and Y position: %d, facing: %s`, r.XPosition, r.YPosition, r.FacingDirection)
}
