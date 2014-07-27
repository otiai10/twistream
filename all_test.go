package twistream_test

// test
import (
	"fmt"
	. "github.com/otiai10/mint"
	"github.com/otiai10/twistream"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

var (
	CONSUMER_KEY        string
	CONSUMER_SECRET     string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
)

func TestNew(t *testing.T) {
	timeline, e := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		CONSUMER_KEY,
		CONSUMER_SECRET,
		ACCESS_TOKEN,
		ACCESS_TOKEN_SECRET,
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

func init() {
	buffer, _ := ioutil.ReadFile("test.conf")
	lines := strings.Split(string(buffer), "\n")
	if len(lines) < 4 {
		panic("test.conf requires at least 4 lines")
	}
	CONSUMER_KEY = lines[0]
	CONSUMER_SECRET = lines[1]
	ACCESS_TOKEN = lines[2]
	ACCESS_TOKEN_SECRET = lines[3]
}
