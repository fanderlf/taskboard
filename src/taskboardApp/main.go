package main

import (
	"html/template"
	"net/http"
	"taskboard"

	"github.com/google/uuid"
)

func init() {
	firstBoard := taskboard.LoadOrNewTaskboard(uuid.New())
	firstBoard.Create("First Team")
	firstBoard.AddStory(uuid.New(), "my first story")
	taskboard.SaveTaskboard(firstBoard)

	secondBoard := taskboard.LoadOrNewTaskboard(uuid.New())
	secondBoard.Create("Second Team")
	secondBoard.AddStory(uuid.New(), "my first story")
	secondBoard.AddStory(uuid.New(), "my second story")
	secondBoard.AddStory(uuid.New(), "my third story")
	taskboard.SaveTaskboard(secondBoard)

	thirdBoard := taskboard.LoadOrNewTaskboard(uuid.New())
	thirdBoard.Create("Third Team")
	thirdBoard.AddStory(uuid.New(), "my first story")
	thirdBoard.AddStory(uuid.New(), "my first story")
	thirdBoard.AddStory(uuid.New(), "my first story")
	thirdBoard.AddStory(uuid.New(), "my first story")
	thirdBoard.AddStory(uuid.New(), "my first story")
	taskboard.SaveTaskboard(thirdBoard)
}

func boards(w http.ResponseWriter, r *http.Request) {
	taskboards := taskboard.GetAllTaskboards()
	mytemplate := `
    <html>
        <head>
            <title>taskboards</title>
        </head>
        <body>
            {{range .}}
                <div>{{.Teamname}} {{.NumberOfStories}}!</div>
            {{end}}
        </body>
    </html>
    `
	t, _ := template.New("boardlist").Parse(mytemplate)
	t.Execute(w, taskboards)
}

func main() {
	http.HandleFunc("/boards", boards)
	http.ListenAndServe(":8080", nil)
}
