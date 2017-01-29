package taskboard

import "github.com/google/uuid"

type TaskboardRead struct {
	Id              uuid.UUID
	Teamname        string
	NumberOfStories int
}

var taskboardRepository = make(map[uuid.UUID]*TaskboardRead)

type StoryRead struct {
	Id          uuid.UUID
	Description string
}

var taskboardStoryRepository = make(map[uuid.UUID][]*StoryRead)

func GetAllTaskboards() []TaskboardRead {
	var boards []TaskboardRead

	for _, value := range taskboardRepository {
		boards = append(boards, *value)
	}

	return boards
}

func GetTaskboardById(taskboardId uuid.UUID) *TaskboardRead {
	return taskboardRepository[taskboardId]
}

func GetStoriesForTaskboard(taskboardId uuid.UUID) []*StoryRead {
	return taskboardStoryRepository[taskboardId]
}

func UpdateReadRepositoryFromEvents(events []interface{}) {
	for _, event := range events {
		switch e := event.(type) {
		case TaskboardCreated:
			taskboardRepository[e.TaskboardId] = &TaskboardRead{
				Id:              e.TaskboardId,
				Teamname:        e.Teamname,
				NumberOfStories: 0,
			}
		case StoryAdded:
			board := taskboardRepository[e.TaskboardId]
			board.NumberOfStories++
			taskboardStoryRepository[e.TaskboardId] = append(taskboardStoryRepository[e.TaskboardId], &StoryRead{
				e.StoryId,
				e.Description,
			})
		}
	}
}
