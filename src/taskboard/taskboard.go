package taskboard

import "github.com/google/uuid"
import "errors"

type Taskstate int

const (
	Todo Taskstate = iota
	Inprogress
	Done
)

type Task struct {
	Id          uuid.UUID
	Description string
	State       Taskstate
}

type Story struct {
	Id          uuid.UUID
	Description string
	Tasks       []Task
}

type Taskboard struct {
	Id               uuid.UUID
	Teamname         string
	Stories          []Story
	initialised      bool
	UncommitedEvents []interface{}
}

type TaskboardInitialised struct {
	TaskboardId uuid.UUID
}

func (board *Taskboard) Init(id uuid.UUID) {
	board.trackChange(TaskboardInitialised{
		id,
	})
}

type TaskboardCreated struct {
	TaskboardId uuid.UUID
	Teamname    string
}

func (board *Taskboard) Create(teamname string) {
	board.trackChange(TaskboardCreated{
		board.Id,
		teamname,
	})
}

type StoryAdded struct {
	StoryId     uuid.UUID
	Description string
	TaskboardId uuid.UUID
}

func (board *Taskboard) AddStory(id uuid.UUID, description string) error {
	if !board.initialised {
		return errors.New("Board has to be created before adding a story")
	}

	board.trackChange(StoryAdded{id, description, board.Id})

	return nil
}

func (board *Taskboard) FromHistory(events []interface{}) *Taskboard {
	for _, event := range events {
		board.transition(event)
	}
	return board
}

func (board *Taskboard) trackChange(event interface{}) {
	board.UncommitedEvents = append(board.UncommitedEvents, event)
	board.transition(event)
}

func (board *Taskboard) transition(event interface{}) {
	switch e := event.(type) {
	case TaskboardInitialised:
		board.Id = e.TaskboardId

	case TaskboardCreated:
		board.Teamname = e.Teamname
		board.initialised = true

	case StoryAdded:
		board.Stories = append(board.Stories, Story{
			Id:          e.StoryId,
			Description: e.Description,
		})
	}
}
