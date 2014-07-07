package twistream_test

import "testing"
import . "github.com/otiai10/mint"
import "github.com/otiai10/twistream"

import "fmt"
import "github.com/robfig/config"
import "time"

func TestNew(t *testing.T) {
	c := getConf()
	timeline, e := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		c["consumer_key"],
		c["consumer_secret"],
		c["access_token"],
		c["access_token_secret"],
	)
	Expect(t, e).ToBe(nil)
	Expect(t, timeline).TypeOf("*twistream.Timeline")

	go func() {
		for {
			status := <-timeline.Listen()
			fmt.Printf("%+v\n", status)
		}
	}()

	time.Sleep(120 * time.Second)
}

func getConf() map[string]string {
	conf, _ := config.ReadDefault("test.conf")
	consumer_key, _ := conf.String("test", "consumer_key")
	consumer_secret, _ := conf.String("test", "consumer_secret")
	access_token, _ := conf.String("test", "access_token")
	access_token_secret, _ := conf.String("test", "access_token_secret")
	return map[string]string{
		"consumer_key":        consumer_key,
		"consumer_secret":     consumer_secret,
		"access_token":        access_token,
		"access_token_secret": access_token_secret,
	}
}
