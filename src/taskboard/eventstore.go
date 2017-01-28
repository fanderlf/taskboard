package taskboard

import (
	"github.com/google/uuid"
)

var store = make(map[uuid.UUID][]interface{})

func LoadOrNewTaskboard(id uuid.UUID) *Taskboard {
	var board = &Taskboard{}

	var events, ok = store[id]

	if ok {
		board.FromHistory(events)
	} else {
		board.Init(id)
	}

	return board
}

func SaveTaskboard(board *Taskboard) {
	eventsToSave := board.UncommitedEvents

	for _, event := range eventsToSave {
		store[board.Id] = append(store[board.Id], event)
	}

	UpdateReadRepositoryFromEvents(eventsToSave)
}
