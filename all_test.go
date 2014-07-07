package twistream_test

import "testing"
import . "github.com/otiai10/mint"
import "github.com/otiai10/twistream"

import "fmt"
import "github.com/robfig/config"

func TestNew(t *testing.T) {
    conf, _ := config.ReadDefault("test.conf")
    consumer_key, _ := conf.String("test", "consumer_key")
    consumer_secret, _ := conf.String("test", "consumer_secret")
    access_token, _ := conf.String("test", "access_token")
    access_token_secret, _ := conf.String("test", "access_token_secret")
	timeline, e := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		consumer_key,
		consumer_secret,
		access_token,
		access_token_secret,
	)
	Expect(t, e).ToBe(nil)
	Expect(t, timeline).TypeOf("*twistream.Timeline")

	fmt.Println(timeline.ReadOnce())
}

/*
func getTimelineOfTest() *twistream.Timeline {
	tl, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		"consumer_key",
		"consumer_secret",
		"access_token",
		"access_token_secret",
	)
	return tl
}
*/
