package twistream

type entities struct {
	Urls         []url     `json:"urls"`
	Hashtags     []hashtag `json:"hashtags"`
	UserMentions []user    `json:"user_mentions"`
}
