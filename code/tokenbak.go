package code

//import (
//	"crypto/hmac"
//	"crypto/sha1"
//	"encoding/base64"
//	"fmt"
//	"net/http"
//	_ "net/http/pprof"
//	"net/url"
//	"strconv"
//	"strings"
//	"time"

//	"github.com/nu7hatch/gouuid"
//)

//const (
//	consumerKey       = "LZNdiWMsZHDkyZd5NJLmMyLsc"
//	consumerKeySecret = "20ee0SSURr3CXKihASbGkHM4FmmQa2xR4GloxKbOP79rgXBcIO"
//	uname             = "Jsonification"
//	paswd             = "891115"
//)

//func GenUUID() string {
//	v4, _ := uuid.NewV4()
//	v5, _ := uuid.NewV5(v4, []byte("twitter"))
//	return v5.String()
//}
//func Tokenb() {
//	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
//	uuid := GenUUID()
//	uuid = strings.Replace(uuid, "-", "", -1)
//	oauth_nonce := base64.StdEncoding.EncodeToString([]byte(uuid))
//	oauth_nonce = strings.Replace(oauth_nonce, "=", "", -1)
//	oauth_nonce = strings.Replace(oauth_nonce, "/", "", -1)
//	oauth_nonce = strings.Replace(oauth_nonce, "+", "", -1)
//	oauth_nonce = strings.Replace(oauth_nonce, "-", "", -1)
//	fmt.Println("oauth_nonce:", oauth_nonce)

//	baseStr := "POST&https%3A%2F%2Fapi.twitter.com%2Foauth%2Faccess_token&" +
//		"oauth_consumer_key%3D" + consumerKey +
//		"%26oauth_nonce%3D" + oauth_nonce +
//		"%26oauth_signature_method%3DHMAC-SHA1" +
//		"%26oauth_timestamp%3D" + timeStr +
//		"%26oauth_version%3D1.0" +
//		"%26x_auth_mode%3Dclient_auth" +
//		"%26x_auth_password%3D" + paswd +
//		"%26x_auth_username%3D" + uname

//	signKey := consumerKeySecret + "&"

//	h := hmac.New(sha1.New, []byte(signKey))
//	h.Write([]byte(baseStr))
//	src := h.Sum(nil)
//	oauth_signature := base64.StdEncoding.EncodeToString(src)

//	vv := url.Values{}
//	vv.Set("oauth_signature", oauth_signature)
//	oauth_signature = strings.Split(vv.Encode(), "=")[1]

//	fmt.Println("oauth_signature:", oauth_signature)

//	baseUrl := "https://api.twitter.com/oauth/access_token"
//	v := url.Values{}
//	v.Set("x_auth_username", uname)
//	v.Set("x_auth_password", paswd)
//	v.Set("x_auth_mode", "client_auth")

//	web := &WebPage{Url: baseUrl, Method: "POST"}
//	web.Proxy = "http://proxy:nMBT@china.opal.me:58892"
//	web.Body = strings.NewReader(v.Encode())

//	authorization := `OAuth oauth_nonce="` + oauth_nonce +
//		`",oauth_signature_method="HMAC-SHA1",oauth_timestamp="` + timeStr +
//		`",oauth_consumer_key="` + consumerKey +
//		`",oauth_signature="` + oauth_signature +
//		`",oauth_version="1.0"`
//	fmt.Println(authorization)
//	header := http.Header{}
//	header.Set("Authorization", authorization)
//	web.Header = header

//	fmt.Println(web.Header)

//	err := web.DoRequest()

//	if err != nil {
//		fmt.Println(err, "--")
//		return
//	}

//	fmt.Println(string(web.RespBody))

//}
