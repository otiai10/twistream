package twistream

import "net/http"
import "github.com/mrjones/oauth"

import "fmt"
import "regexp"
import "encoding/json"

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
	response, e := consumer.Get(
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

type parser struct {
	streamProxy chan Status
}

// TODO: make it not global
var pool = []byte{}
var flag = false

// TODO: refactoring
func (p parser) Write(message []byte) (n int, err error) {
	if flag {
		pool = append(pool, message...)
		initial := map[string][]int{}
		status := Status{}
		if err = json.Unmarshal(pool, &initial); err == nil {
			if _, ok := initial["friends"]; ok {
				// Do nothing for initial entry
			}
			pool = []byte{}
		} else if err = json.Unmarshal(pool, &status); err == nil {
			p.streamProxy <- status
			pool = []byte{}
		}
	}
	flag = false
	if regexp.MustCompile("^[0-9a-z]+\r\n$").Match(message) {
		// Just before entity
		flag = true
	}
	return
}

// Listen provides channel of Tweet.
func (tl *Timeline) Listen() chan Status {
	p := parser{
		tl.stream,
	}
	go func() {
		for {
			tl.response.Write(p)
		}
	}()
	return tl.stream
}
