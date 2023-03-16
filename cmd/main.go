package main

import (
	"fmt"

	"github.com/satori/go.uuid"

	//_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	task "github.com/LepikovStan/plagiarismchecker/task"
)

// type Task struct {
// 	ID              string
// 	State           string
// 	OriginalArticle string
// 	ErrorMessage    string
// }

func main() {
	taskID := uuid.NewV4().String()
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }
	t := task.Task{ID: taskID, State: "created", OriginalArticle: "d.Message"}
	fmt.Println(t)
}
