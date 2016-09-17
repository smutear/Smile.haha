package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	//	"smile.haha/code"
)

func main() {
	//	tc := code.TwitterClient{}
	//	tc.Main()

	engin := gin.New()

	store := sessions.NewCookieStore([]byte("secret"))
	options := sessions.Options{MaxAge: 10}
	store.Options(options)
	engin.Use(sessions.Sessions("GINSESSIONID", store))

	//	engin.Use(func(c *gin.Context) {
	//		sn := sessions.Default(c)
	//		key := sn.Flashes()
	//		fmt.Println("flush:", key)
	//	})

	engin.GET("/login", func(c *gin.Context) {
		user := c.Query("user")
		sn := sessions.Default(c)
		sn.Set("user", user)
		sn.Save()
		sn.AddFlash(user)
		key := sn.Flashes()
		fmt.Println("flush:", key)
		c.JSON(http.StatusOK, "ok")
	})
	engin.GET("/show", func(c *gin.Context) {
		sn := sessions.Default(c)
		user := sn.Get("user")
		if user == nil {
			c.JSON(http.StatusOK, "not login")
			return
		}
		key := sn.Flashes()
		fmt.Println("flush:", key)
		c.JSON(http.StatusOK, "login:"+user.(string))
	})
	engin.Run(":9000")
}
