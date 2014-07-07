package twistream

type url struct {
	Url         string `json:"url"`
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Indices     []int  `json:"indices"`
}
