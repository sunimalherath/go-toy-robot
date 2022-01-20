package util

func IsValidPlacement(x, y int) bool {
	// check if the robot is placed within the 5 x 5 grid
	if x < 0 || x > 4 {
		return false
	}
	if y < 0 || y > 4 {
		return false
	}
	return true
}

func IsValidCommand(cmd command) bool {
	validCommands := []command{PLACE, MOVE, LEFT, RIGHT, REPORT}
	for _, v := range validCommands {
		if cmd == v {
			return true
		}
	}
	return false
}

func IsValidMove(x, y int, facingDirection Direction) bool {
	switch facingDirection {
	case NORTH:
		return IsValidPlacement(x, y+1)
	case SOUTH:
		return IsValidPlacement(x, y-1)
	case EAST:
		return IsValidPlacement(x+1, y)
	case WEST:
		return IsValidPlacement(x-1, y)
	default:
		return false
	}
}
