package page

type Page struct {
	ID                 int    `json:"id"`
	PlagiarisedPercent int    `json:"plagiarised_percent"`
	ReportID           int    `json:"report_id"`
	Title              string `json:"title"`
	URL                string `json:"url"`
	Description        string `json:"description"`
}
