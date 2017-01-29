package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"taskboard"

	"fmt"

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
			<script src="/file?name=/js/test.js"></script>
        </head>
        <body>
            {{range .}}<div>{{.Teamname}} {{.NumberOfStories}}!</div>
			{{end}}
        </body>
    </html>
    `
	t, _ := template.New("boardlist").Parse(mytemplate)
	t.Execute(w, taskboards)
}

func file(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("name")

	if filename == "" {
		http.Error(w, "you have to provide a filename", http.StatusNotFound)
		return
	}

	filepath := "./assets" + filename

	fmt.Println(filepath)

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprintf(w, "%s", content)
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/boards", boards)
	http.HandleFunc("/file", file)
	http.ListenAndServe(":8080", nil)
}
