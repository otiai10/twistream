// Definitions of properties in Status
package twistream

type entities struct {
	Urls         []url     `json:"urls"`
	Hashtags     []hashtag `json:"hashtags"`
	UserMentions []user    `json:"user_mentions"`
}
type geo struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}
type hashtag struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
}
type url struct {
	Url         string `json:"url"`
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Indices     []int  `json:"indices"`
}
type user struct {
	Id                             int64  `json:"id"`
	IdStr                          string `json:"id_str"`
	Name                           string `json:"name"`
	ScreenName                     string `json:"screen_name"`
	Location                       string `json:"location"`
	Url                            string `json:"url"`
	Description                    string `json:"string"`
	Protected                      bool   `json:"protected"`
	FollowersCount                 int    `json:"followers_count"`
	FriendsCount                   int    `json:"friends_count"`
	ListedCount                    int    `json:"listed_count"`
	CreatedAt                      string `json:"created_at"`
	FavouritesCount                int    `json:"favourites_count"`
	UtcOffset                      int    `json:"utc_offset"`
	TimeZone                       string `json:"time_zone"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	Verified                       bool   `json:"verified"`
	StatusesCount                  int64  `json:"statuses_count"`
	Lang                           string `json:"lang"`
	ContributorsEnabled            bool   `json:"contributors_enabled"`
	IsTranslator                   bool   `json:"is_translator"`
	IsTranslationEnabled           bool   `json:"is_translation_enabled"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageUrl      string `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileImageUrl                string `json:"profile_image_url"`
	ProfileImageUrlHttps           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	DefaultProfile                 bool   `json:"default_profile"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	Following                      bool   `json:"following"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	Notifications                  bool   `json:"notifications"`
}
