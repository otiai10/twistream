package twistream

type Timeline struct{}

// New provides new reference for specified Timeline.
func New(endpoint, consumerKey, consumerSecret, accessToken, accessTokenSecret string) (tl *Timeline, e error) {
	return
}

// Listen provides channel of Tweet.
func (tl *Timeline) Listen(fltr ...interface{}) (ch *chan<- Status) {
	return
}
