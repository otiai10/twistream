package twistream_test

import "github.com/otiai10/twistream"
import "testing"
import "fmt"

func ExampleTimeline_Listen(t *testing.T) {
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		"CONSUMER_KEY",
		"CONSUMER_SECRET",
		"ACCESS_TOKEN",
		"ACCESS_TOKEN_SECRET",
	)
	// Practically, you would run this loop by goroutine
	for {
		status := <-timeline.Listen()
		fmt.Println(status)
	}
}
