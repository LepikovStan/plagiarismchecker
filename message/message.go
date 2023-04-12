package message

import (
	"github.com/LepikovStan/plagiarismchecker/report"
	"github.com/LepikovStan/plagiarismchecker/task"
)

type Data struct {
	Task   task.Task
	Report report.Report
}
