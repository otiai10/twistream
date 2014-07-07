package twistream

type Entities struct {
	Urls         []Url     `json:"urls"`
	Hashtags     []Hashtag `json:"hashtags"`
	UserMentions []User    `json:"user_mentions"`
}
