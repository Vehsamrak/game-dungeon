package app

type RoomMemoryRepository struct {
	rooms []*Room
}

func (repository RoomMemoryRepository) Create(rooms []*Room) RoomRepository {
	if rooms == nil {
		rooms = []*Room{
			{x: -1, y: 0},
			{x: 1, y: 1},
		}
	}

	return &RoomMemoryRepository{rooms: rooms}
}

func (repository *RoomMemoryRepository) FindByXandY(x int, y int) *Room {
	for _, room := range repository.rooms {
		if room.x == x && room.y == y {
			return room
		}
	}

	return nil
}

func (repository *RoomMemoryRepository) FindByXY(XY XYInterface) *Room {
	return repository.FindByXandY(XY.X(), XY.Y())
}

func (repository *RoomMemoryRepository) AddRoom(room *Room) {
	// TODO: lock
	repository.rooms = append(repository.rooms, room)
	// TODO: unlock
}
