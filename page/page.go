package page

import (
	"github.com/lib/pq"
)

type Page struct {
	ID                   int           `json:"id"`
	PlagiarisedPercent   int           `json:"plagiarised_percent"`
	ReportID             int           `json:"report_id"`
	Title                string        `json:"title"`
	URL                  string        `json:"url"`
	Description          string        `json:"description"`
	PlagiarisedSentences pq.Int64Array `gorm:"type:integer[]"`
}
