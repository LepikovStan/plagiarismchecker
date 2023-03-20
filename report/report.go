package report

import "time"

type Report struct {
	ID                   int
	UniqueSentences      int
	PlagiarisedPercent   int
	AIGeneratedPercent   int
	ReadabilityPercent   int
	ErrorsCount          int
	UniquePercent        int
	TotalWords           int
	TotalChars           int
	PlagiarisedSentences []int
	PlagiarisedPages     []int
	TaskID               string
	CreatedAt            time.Time
}

func NewPageReport() {}
