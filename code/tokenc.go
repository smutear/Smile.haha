package code

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nu7hatch/gouuid"
)

const (
	//	consumerKey       = "xvz1evFS4wEEPTGEFPHBog"
	//	consumerKeySecret = "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw"
	//	accessToken       = "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"
	//	accessTokenSecret = "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"

	consumerKey       = "LZNdiWMsZHDkyZd5NJLmMyLsc"
	consumerKeySecret = "20ee0SSURr3CXKihASbGkHM4FmmQa2xR4GloxKbOP79rgXBcIO"
	accessToken       = "732495514009116673-8d69TSSAEFs0ezhRDf0h6RJcQ8aqV3W"
	accessTokenSecret = "Cup12D85vnrKDgMdhWgWfeRFCDLiFKtdEhYgC1vBgSihI"
)

func GenUUID() string {
	v4, _ := uuid.NewV4()
	v5, _ := uuid.NewV5(v4, []byte("twitter"))
	return v5.String()
}
func Tokenb() {
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	//	timeStr = "1318622958"
	uuid := GenUUID()
	uuid = strings.Replace(uuid, "-", "", -1)
	oauth_nonce := base64.StdEncoding.EncodeToString([]byte(uuid))
	oauth_nonce = strings.Replace(oauth_nonce, "=", "", -1)
	oauth_nonce = strings.Replace(oauth_nonce, "/", "", -1)
	oauth_nonce = strings.Replace(oauth_nonce, "+", "", -1)
	oauth_nonce = strings.Replace(oauth_nonce, "-", "", -1)
	//	oauth_nonce = "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"
	fmt.Println("oauth_nonce:", oauth_nonce)

	baseStr := "POST&https%3A%2F%2Fapi.twitter.com%2F1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26" +
		"oauth_consumer_key%3D" + consumerKey +
		"%26oauth_nonce%3D" + oauth_nonce +
		"%26oauth_signature_method%3DHMAC-SHA1" +
		"%26oauth_timestamp%3D" + timeStr +
		"%26oauth_token%3D" + accessToken +
		"%26oauth_version%3D1.0" +
		"%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"

	signKey := consumerKeySecret + "&" + accessTokenSecret

	h := hmac.New(sha1.New, []byte(signKey))
	h.Write([]byte(baseStr))
	src := h.Sum(nil)
	oauth_signature := base64.StdEncoding.EncodeToString(src)

	vv := url.Values{}
	vv.Set("oauth_signature", oauth_signature)
	oauth_signature = strings.Split(vv.Encode(), "=")[1]

	fmt.Println("oauth_signature:", oauth_signature)

	baseUrl := "https://api.twitter.com/oauth/request_token"
	v := url.Values{}
	v.Set("oauth_callback", "http://dudu.city/")

	web := &WebPage{Url: baseUrl, Method: "POST"}
	web.Proxy = "http://proxy:nMBT@china.opal.me:58892"
	web.Body = strings.NewReader(v.Encode())

	authorization := `OAuth ` + v.Encode() + `, oauth_nonce="` + oauth_nonce +
		`", oauth_signature_method="HMAC-SHA1", oauth_timestamp="` + timeStr +
		`", oauth_consumer_key="` + consumerKey +
		`", oauth_token="` + accessToken +
		`", oauth_signature="` + oauth_signature +
		`", oauth_version="1.0"`

	//	authorization := "OAuth oauth_nonce=\"" + oauth_nonce + "\", oauth_callback=\"http%3A%2F%2Fmyapp.com%3A3005%2Ftwitter%2Fprocess_callback\", oauth_signature_method=\"HMAC-SHA1\", oauth_timestamp=\"" + timeStr + "\", oauth_consumer_key=\"" + consumerKey + "\", oauth_signature=\"" + oauth_signature + "\", oauth_version=\"1.0\""
	fmt.Println(authorization)
	header := http.Header{}
	header.Set("Authorization", authorization)
	web.Header = header

	err := web.DoRequest()

	if err != nil {
		fmt.Println(err, "--")
		return
	}

	fmt.Println(web.Status, web.StatusCode)
	fmt.Println(string(web.RespBody))

}
