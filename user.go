package twistream

type User struct {
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

/*
"user":{"id":971441053,"id_str":"971441053","name":"hisyotan","screen_name":"hisyotan","location":"","url":"http:\/\/otiai10.com\/hisyotan\/","description":"hisyotan\u306f\u3057\u3070\u3089\u304f\u65c5\u306b\u51fa\u307e\u3059\u3002\u63a2\u3055\u306a\u3044\u3067\u304f\u3060\u3055\u3044","protected":false,"followers_count":14,"friends_count":2,"listed_count":0,"created_at":"Mon Nov 26 06:24:02 +0000 2012","favourites_count":2,"utc_offset":32400,"time_zone":"Irkutsk","geo_enabled":false,"verified":false,"statuses_count":1935,"lang":"ja","contributors_enabled":false,"is_translator":false,"is_translation_enabled":false,"profile_background_color":"C0DEED","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/378800000178498798\/2a1604598ab1897d0485aae57c278733_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/378800000178498798\/2a1604598ab1897d0485aae57c278733_normal.png","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"default_profile":true,"default_profile_image":false,"following":null,"follow_request_sent":null,"notifications":null}
*/
