package app_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/game-dungeon/internal/app"
	"testing"
)

func TestRoom(test *testing.T) {
	suite.Run(test, &roomTest{})
}

type roomTest struct {
	suite.Suite
}

func (suite *roomTest) Test_Create_roomParameters_newRoomCreated() {
	x, y := 1, 1

	room := app.Room{}.Create(x, y)

	assert.Equal(suite.T(), x, room.X())
	assert.Equal(suite.T(), y, room.Y())
}

func (suite *roomTest) Test_AddFlag_roomWithoutFlags_flagsAddedToRoom() {
	room := suite.createRoom()
	firstFlag := "first_flag"
	secondFlag := "second_flag"
	flags := []string{secondFlag}

	room.AddFlag(firstFlag)
	room.AddFlags(flags)

	assert.True(suite.T(), room.HasFlag(firstFlag))
	assert.True(suite.T(), room.HasFlag(secondFlag))
}

func (suite *roomTest) createRoom() *app.Room {
	return app.Room{}.Create(0, 0)
}