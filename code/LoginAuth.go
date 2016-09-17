package code

//import (
//	"fmt"
//	"log"
//	"net/http"

//	"github.com/gin-gonic/contrib/sessions"
//	"github.com/gin-gonic/gin"
//)

//type LoginAuth struct {
//}

//func (la *LoginAuth) Main() {
//	r := gin.Default()
//	//	store := sessions.NewCookieStore([]byte("secret"))
//	store, err := sessions.NewRedisStore(10, "tcp", "192.168.1.103:6379", "", []byte("secret"))
//	if err != nil {
//		panic(err)
//	}
//	options := sessions.Options{MaxAge: 10}
//	store.Options(options)
//	r.Use(sessions.Sessions("user-auth", store))

//	authorized := r.Group("/")
//	authorized.Use(la.auth)
//	{
//		authorized.GET("/ping", func(c *gin.Context) {
//			c.JSON(http.StatusOK, "ok")
//		})
//		authorized.GET("/ops", func(c *gin.Context) {
//			c.JSON(http.StatusOK, "ops")
//		})
//	}
//	login := r.Group("/")
//	{
//		login.GET("/index", func(c *gin.Context) {
//			c.JSON(http.StatusOK, "this is index page for login.")
//		})
//		login.GET("/login", func(c *gin.Context) {
//			sn := sessions.Default(c)
//			fmt.Println("sn in login:")
//			fmt.Println(sn)
//			sn.Set("user", "a")
//			err := sn.Save()
//			if err != nil {
//				log.Fatal(err)
//				c.JSON(http.StatusInternalServerError, "session save failed.")
//			}
//			c.JSON(http.StatusOK, "login seccess.")
//		})
//		login.GET("/clogin", func(c *gin.Context) {
//			c.SetCookie("demo.user_id", "random-id-01", 10, "", "", false, false)
//			c.JSON(http.StatusOK, "c login success.")
//		})
//	}

//	cookieLogin := r.Group("/")
//	cookieLogin.Use(la.cookieAuth)
//	{
//		cookieLogin.GET("/cops", func(c *gin.Context) {
//			c.JSON(http.StatusOK, "cops success.")
//		})
//	}

//	r.Run(":9000")
//}

//func (la *LoginAuth) cookieAuth(c *gin.Context) {
//	user, err := c.GetCookie("demo.user_id")
//	//	if  {
//	//		log.Fatal(err)
//	//		c.JSON(http.StatusInternalServerError, "error")
//	//		return
//	//	}
//	if err != nil || user == "" {
//		fmt.Println(err)
//		c.Redirect(http.StatusFound, "/index")
//		return
//	}
//	c.SetCookie("demo.user_id", "random-id-01", 10, "", "", false, false)
//}

//func (la *LoginAuth) auth(c *gin.Context) {
//	sn := sessions.Default(c)
//	user := sn.Get("user")
//	if user == nil {
//		c.Redirect(http.StatusFound, "/index")
//		return
//	}
//	//too many session instances.
//	//	sn.Set("user", user)
//	//	sn.Save()
//	fmt.Println(sn)
//}
