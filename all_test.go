package twistream_test

// test
import (
	"bufio"
	"fmt"
	. "github.com/otiai10/mint"
	"github.com/otiai10/twistream"
	"os"
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
	time.Sleep(60 * time.Second)
}

func init() {
	f, _ := os.Open("conf")
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	CONSUMER_KEY = s.Text()
	s.Scan()
	CONSUMER_SECRET = s.Text()
	s.Scan()
	ACCESS_TOKEN = s.Text()
	s.Scan()
	ACCESS_TOKEN_SECRET = s.Text()
}
