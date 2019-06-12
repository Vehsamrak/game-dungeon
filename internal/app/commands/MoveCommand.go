package commands

import (
	"github.com/vehsamrak/game-dungeon/internal/app"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/direction"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/gameError"
	"github.com/vehsamrak/game-dungeon/internal/app/enum/roomFlag"
)

type MoveCommand struct {
	roomRepository app.RoomRepository
}

func (command MoveCommand) Create(roomRepository app.RoomRepository) *MoveCommand {
	return &MoveCommand{roomRepository: roomRepository}
}

func (command *MoveCommand) Execute(character Character, arguments ...string) (result CommandResult) {
	result = commandResult{}.Create()

	if len(arguments) < 1 {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	moveDirection, err := direction.FromString(arguments[0])
	if err != "" {
		result.AddError(gameError.WrongCommandAttributes)
		return
	}

	xDiff, yDiff, zDiff := moveDirection.DiffXYZ()
	x := character.X() + xDiff
	y := character.Y() + yDiff
	z := character.Z() + zDiff

	room := command.roomRepository.FindByXYandZ(x, y, z)

	err = command.checkRoomMobility(room)
	if err == "" {
		character.Move(x, y, z)
	} else {
		result.AddError(err)
	}

	return
}

func (command *MoveCommand) checkRoomMobility(room *app.Room) (err gameError.Error) {
	if room == nil {
		err = gameError.RoomNotFound
	} else if room.HasFlag(roomFlag.Unfordable) {
		err = gameError.RoomUnfordable
	}

	return
}
