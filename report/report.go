package report

type Report struct {
	ID                   int
	UniqueSentences      int
	PlagiarisedPercent   int
	UniquePercent        int
	TotalWords           int
	TotalChars           int
	PlagiarisedSentences []int
	PlagiarisedPages     []int
	TaskID               string
}
