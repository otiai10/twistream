package twistream

import "log"
import "net/http"
import "bufio"
import "encoding/json"

type Timeline struct {
	response *http.Response
	stream   chan Status
	client   *api
}

// New provides new reference for specified Timeline.
func New(endpoint, consumerKey, consumerSecret, accessToken, accessTokenSecret string) (tl *Timeline, e error) {

	tl = &Timeline{
		client: initAPI(
			consumerKey,
			consumerSecret,
			accessToken,
			accessTokenSecret,
		),
	}

	response, e := tl.client.Get(
		endpoint,
		map[string]string{},
	)
	tl.response = response
	tl.stream = make(chan Status)
	return
}

// Listen bytes sent from Twitter Streaming API
// and send completed status to the channel.
func (tl *Timeline) Listen() <-chan Status {
	scanner := bufio.NewScanner(tl.response.Body)
	go func() {
		for {
			if ok := scanner.Scan(); !ok {
				continue
			}
			status := new(Status)
			if err := json.Unmarshal(scanner.Bytes(), &status); err != nil {
				log.Println("(abort)")
				continue
			}
			tl.stream <- *status
		}
	}()
	return tl.stream
}

// Tweet posts status to the timeline
func (tl *Timeline) Tweet(status Status) (e error) {
	_, e = tl.client.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		status.ToParams(),
	)
	return
}
