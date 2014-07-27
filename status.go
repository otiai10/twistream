package twistream

import "strconv"

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

func (s Status) ToParams() map[string]string {
	// TODO: embed more information
	params := map[string]string{
		"status": s.Text,
	}
	if s.InReplyToStatusId != 0 {
		params["in_reply_to_status_id"] = strconv.FormatInt(s.InReplyToStatusId, 10)
	}
	return params
}
