# Twitter Streaming API [![GoDoc](https://godoc.org/github.com/otiai10/twistream?status.svg)](https://godoc.org/github.com/otiai10/twistream)

The very simplest interface to use Twitter Streaming API by golang.

# Usage

```go
timeline, _ := twistream.New(
    "https://userstream.twitter.com/1.1/user.json",
    CONSUMERKEY,
    CONSUMERSECRET,
    ACCESSTOKEN,
    ACCESSTOKENSECRET,
)

// Listen timeline
for {
    status := <-timeline.Listen()
    fmt.Println(status)
}

// Tweet to timeline
status := twistream.Status{
    Text: "@otiai10 How does Go work?",
    InReplyToStatusId: 493324823926288386,
}
_ := timeline.Tweet(status)
```

# TODOs

- [x] GET user
- [ ] GET site
- [ ] GET statuses/sample
- [ ] GET status/firehose
- [ ] POST statuses/filter
- [x] POST statuses/update
- [ ] POST statuses/update_with_media
