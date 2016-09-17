package code

import (
	"fmt"
	"net/url"

	"database/sql"

	"github.com/ChimeraCoder/anaconda"
	//	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type TwitterClient struct {
}

func (tc *TwitterClient) Main() {
	anaconda.SetConsumerKey("LZNdiWMsZHDkyZd5NJLmMyLsc")
	anaconda.SetConsumerSecret("20ee0SSURr3CXKihASbGkHM4FmmQa2xR4GloxKbOP79rgXBcIO")
	api := anaconda.NewTwitterApi("732495514009116673-Xwy1Uc2tHXSH116iCSpQ9BOnNZuc2x2", "EzvrsZXNQsX9wXBnRH88dx2CibCsdohBHhADczwaPAv9J")
	v := url.Values{}
	v.Set("screen_name", "Jsonification")
	res, _ := api.GetUserTimeline(v)

	db, _ := sql.Open("mysql", uri())

	for _, v := range res {
		fmt.Println(v.Text)
		_, err := db.Exec("insert into xg_tmp values(?)", v.Text)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func uri() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", "datamesh", "Hello2016", "47.90.38.89", "3306", "twitter_undp", "utf8")
}
