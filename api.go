package twistream

import "github.com/mrjones/oauth"
import "net/http"

type api struct {
	consumer *oauth.Consumer
	token    *oauth.AccessToken
}

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

func initAPI(ck, cs, at, as string) *api {
	consumer := oauth.NewConsumer(
		ck,
		cs,
		provider,
	)
	token := &oauth.AccessToken{
		at,
		as,
		make(map[string]string),
	}
	return &api{
		consumer,
		token,
	}
}

func (a *api) DispatchRequestFunc(endpoint string) func(string, map[string]string) (*http.Response, error) {
	return a.Get
}

func (a *api) Get(endpoint string, params map[string]string) (*http.Response, error) {
	return a.consumer.Get(
		endpoint,
		params,
		a.token,
	)
}
func (a *api) Post(endpoint string, params map[string]string) (*http.Response, error) {
	return a.consumer.Post(
		endpoint,
		params,
		a.token,
	)
}
