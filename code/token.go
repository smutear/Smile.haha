package code

//import (
//	//	"bytes"
//	"fmt"
//	"net/http"
//	"net/url"
//	"strings"
//)

//const (
//	consumerKey       = "JvyS7DO2qd6NNTsXJ4E7zA"
//	consumerKeySecret = "9z6157pUbOBqtbm0A0q4r29Y2EYzIHlUwbF4Cl9c"
//	uname             = "oauth_test_exec"
//	paswd             = "twitter-xauth"
//)

//func Token() {
//	baseUrl := "https://api.twitter.com/oauth/access_token"
//	v := url.Values{}
//	v.Set("x_auth_username", uname)
//	v.Set("x_auth_password", paswd)
//	v.Set("x_auth_mode", "client_auth")

//	web := &WebPage{Url: baseUrl, Method: "POST"}
//	web.Body = strings.NewReader(v.Encode())
//	web.Proxy = "http://proxy:nMBT@china.opal.me:58892"

//	authorization := `OAuth oauth_nonce="6AN2dKRzxyGhmIXUKSmp1JcB4pckM8rD3frKMTmVAo", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1284565601", oauth_consumer_key="JvyS7DO2qd6NNTsXJ4E7zA", oauth_signature="1L1oXQmawZAkQ47FHLwcOV%2Bkjwc%3D", oauth_version="1.0"`
//	header := http.Header{}
//	header.Set("Authorization", authorization)
//	//	header.Set("Content-Type", "application/x-www-form-urlencoded")
//	web.Header = header

//	err := web.DoRequest()

//	if err != nil {
//		fmt.Println(err, "--")
//		return
//	}

//	fmt.Println(string(web.RespBody))
//}
