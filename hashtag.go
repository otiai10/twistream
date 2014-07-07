package twistream

type Hashtag struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
	// UserMentions []
}
