package util

type command string

const (
	PLACE  command = "place"
	MOVE   command = "move"
	LEFT   command = "left"
	RIGHT  command = "right"
	REPORT command = "report"
)

type Direction string

const (
	NORTH Direction = "north"
	SOUTH Direction = "south"
	EAST  Direction = "east"
	WEST  Direction = "west"
)
