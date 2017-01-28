package taskboard

import "testing"
import "github.com/google/uuid"

func TestTaskboardInitialise(t *testing.T) {
	t.Log("When initialising a new taskboard")
	id := uuid.New()
	board := LoadOrNewTaskboard(id)

	board.Create("First Team")

	t.Log(" it should set the teamname")
	if board.Teamname != "First Team" {
		t.Errorf("Team name was not set correctly. Expected: '%s' Actually:'%s", "First Team", board.Teamname)
	}

	t.Log(" it should set the id")
	if board.Id.String() != id.String() {
		t.Errorf("Id was not set correctly. Expected: '%s' Actually:'%s", id.String(), board.Id.String())
	}
}

func TestTaskboardAddStory(t *testing.T) {
	t.Log("When adding a story to a taskboard")
	id := uuid.New()
	board := LoadOrNewTaskboard(id)

	board.Create("First Team")

	storyId := uuid.New()

	board.AddStory(storyId, "Story1")

	t.Log(" it should add the story")
	if story := board.Stories[0]; story.Id.String() != storyId.String() || story.Description != "Story1" || len(board.Stories) != 1 {
		t.Errorf("Story was not added successfull")
	}
}
