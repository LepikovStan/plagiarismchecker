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
	PlagiarisedSentences int64 `gorm:"type:integer[]"`
	TaskID               string
	CreatedAt            time.Time
}

func NewPageReport() {}
