package twistream

type Status struct {
	CreatedAt            string   `json:"created_at"`
	Id                   int64    `json:"id"`
	IdStr                string   `json:"id_str"`
	Text                 string   `json:"text"`
	Source               string   `json:"source"`
	Truncated            bool     `json:"truncated"`
	InReplyToStatusId    int64    `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string   `json:"in_reply_to_status_id_str"`
	InReplyToUserId      int64    `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string   `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string   `json:"in_reply_to_screen_name"`
	User                 user     `json:"user"`
	Geo                  geo      `json:"geo"`
	Place                string   `json:"place"`
	Contributors         []int64  `json:"contributors"`
	RetweetCount         int      `json:"retweet_count"`
	FavoriteCount        int      `json:"favorite_count"`
	Entities             entities `json:"entities"`
	Favorited            bool     `json:"favorited"`
	Retweeted            bool     `json:"retweeted"`
	FilterLevel          string   `json:"filter_level"`
	Lang                 string   `json:"lang"`
}
