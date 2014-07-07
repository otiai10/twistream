package twistream

type hashtag struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
	// UserMentions []
}
