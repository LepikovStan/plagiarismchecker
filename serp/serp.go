package serp

type SERP struct {
	StatusCode int8     `json:"statusCode"`
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
	URL                    string `json:"url"`
	NumberOfResults        int16  `json:"number_of_results"`
	Location               string `json:"location"`
	NumberOfOrganicResults int16  `json:"number_of_organic_results"`
	NumberOfAds            int8   `json:"number_of_ads"`
	NumberOfPage           int8   `json:"number_of_page"`
	NoResultspage          string `json:"no_results_message"`
}
