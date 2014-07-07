# Twitter Streaming API

twistream

# Usage

```go
timeline := twistream.New(
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
