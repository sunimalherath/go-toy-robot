package util

type IValidator interface {
	IsValidPlacement(int, int) bool
	IsValidCommand(command) bool
	IsValidMove(int, int, Direction) bool
}

type SupportCommand struct{}

func (s SupportCommand) IsValidPlacement(x, y int) bool {
	// check if the robot is placed within the 5 x 5 grid
	if x < 0 || x > 4 {
		return false
	}
	if y < 0 || y > 4 {
		return false
	}
	return true
}

func (s SupportCommand) IsValidCommand(cmd command) bool {
	validCommands := []command{PLACE, MOVE, LEFT, RIGHT, REPORT}
	for _, v := range validCommands {
		if cmd == v {
			return true
		}
	}
	return false
}

func (s SupportCommand) IsValidMove(x, y int, facingDirection Direction) bool {
	switch facingDirection {
	case NORTH:
		return s.IsValidPlacement(x, y+1)
	case SOUTH:
		return s.IsValidPlacement(x, y-1)
	case EAST:
		return s.IsValidPlacement(x+1, y)
	case WEST:
		return s.IsValidPlacement(x-1, y)
	default:
		return false
	}
}

func NewSupportCommand() IValidator {
	return &SupportCommand{}
}
