package twistream

import "github.com/mrjones/oauth"
import "net/http"
import "regexp"

type Timeline struct {
	response *http.Response
	stream   chan Status
}

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

// New provides new reference for specified Timeline.
func New(endpoint, consumerKey, consumerSecret, accessToken, accessTokenSecret string) (tl *Timeline, e error) {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		provider,
	)
	token := &oauth.AccessToken{
		accessToken,
		accessTokenSecret,
		make(map[string]string),
	}
	request := dispatchRequestMethod(endpoint, consumer)
	response, e := request(
		endpoint,
		map[string]string{},
		token,
	)
	tl = &Timeline{
		response,
		make(chan Status),
	}
	return
}

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() <-chan Status {
	// Delegate channel to parser.
	p := &parser{
		proxy:   tl.stream,
		trigger: regexp.MustCompile("^[0-9a-z]+\r\n$"),
	}
	go func() {
		for {
			tl.response.Write(p)
		}
	}()
	return tl.stream
}

// dispatchRequestMethod decides method to request.
// Returns `Get` or `Post`
func dispatchRequestMethod(endpoint string, consumer *oauth.Consumer) func(string, map[string]string, *oauth.AccessToken) (*http.Response, error) {
	return consumer.Get
}
