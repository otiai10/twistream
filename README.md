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

for {
    status := <-timeline.Listen()
    fmt.Println(status)
}
```
