package twistream_test

import "testing"
import . "github.com/otiai10/mint"
import "github.com/otiai10/twistream"

func TestNew(t *testing.T) {
	timeline, e := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		"consumer_key",
		"consumer_secret",
		"access_token",
		"access_token_secret",
	)
	Expect(t, e).ToBe(nil)
	Expect(t, timeline).TypeOf("*twistream.Timeline")
}

func TestTimeline_Listen(t *testing.T) {
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		"consumer_key",
		"consumer_secret",
		"access_token",
		"access_token_secret",
	)
	Expect(t, timeline.Listen()).TypeOf("*chan<- twistream.Status")
}
