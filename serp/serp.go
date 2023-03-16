package serp

type SERP struct {
	StatusCode int      `json:"statusCode"`
	Body       SERPBody `json:"body"`
}

type SERPBody struct {
	MetaData       SERPMetaData        `json:"meta_data"`
	OrganicResults []SERPOrganicResult `json:"organic_results"`
}

type SERPOrganicResult struct {
	URL string `json:"url"`
}

type SERPMetaData struct {
	NumberOfResults        int    `json:"number_of_results"`
	NumberOfOrganicResults int    `json:"number_of_organic_results"`
	NumberOfAds            int    `json:"number_of_ads"`
	NumberOfPage           int    `json:"number_of_page"`
	URL                    string `json:"url"`
	Location               string `json:"location"`
	NoResultspage          string `json:"no_results_message"`
}
