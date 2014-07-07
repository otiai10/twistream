package twistream

import "net/http"
import "github.com/mrjones/oauth"

import "fmt"
import "regexp"
import "encoding/json"

type Timeline struct {
	response *http.Response
}

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

// New provides new reference for specified Timeline.
func New(endpoint, consumerKey, consumerSecret, accessToken, accessTokenSecret string) (tl *Timeline, e error) {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		provider,
	)
	token := &oauth.AccessToken{
		accessToken,
		accessTokenSecret,
		make(map[string]string),
	}
	response, e := consumer.Get(
		endpoint,
		map[string]string{},
		token,
	)
	tl = &Timeline{response}
	return
}

// Listen provides channel of Tweet.
func (tl *Timeline) Listen(fltr ...interface{}) (ch *chan<- Status) {
	return
}

type parser struct {
}

var pool = []byte{}
var flag = false

func (p parser) Write(message []byte) (n int, err error) {
	if flag {
		pool = append(pool, message...)
		initial := map[string][]int{}
		status := Status{}
		if err = json.Unmarshal(pool, &initial); err == nil {
			if _, ok := initial["friends"]; ok {
				fmt.Println("これイニシャルだよー. 却下")
			}
			pool = []byte{}
		} else if err = json.Unmarshal(pool, &status); err == nil {
			fmt.Printf("statusでUnmarshal成功!!\n-------------\n%+v\n------------------\n%v\n", status, string(pool))
			pool = []byte{}
		}
	}
	flag = false
	if regexp.MustCompile("^[0-9a-z]+\r\n$").Match(message) {
		// これはエンティティの一個前
		flag = true
	}
	return
}

func (tl *Timeline) ReadOnce() string {
	p := parser{}
	for {
		// こっから別
		e := tl.response.Write(p)
		fmt.Println("ERROR:", e)
	}
	return "hoge"
}
