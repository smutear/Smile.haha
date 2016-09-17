package code

//import (
//	"fmt"

//	"github.com/gin-gonic/gin"
//)

//type FileUpload struct {
//}

//func (fu *FileUpload) Main() {
//	r := gin.Default()
//	r.POST("/upload", fu.upload)
//	r.Run(":8000")
//}

//func (fu *FileUpload) upload(c *gin.Context) {
//	_, header, _ := c.Request.FormFile("upload")
//	fmt.Println(header.Filename)
//	fmt.Println("hh")
//	fmt.Println(c.Request.MultipartForm)
//	if c.Request.MultipartForm != nil && c.Request.MultipartForm.File != nil {
//		values := c.Request.MultipartForm.File["upload"] //.Value["upload"]
//		fmt.Println(len(values))
//		if len(values) > 0 {
//			for k, v := range values {
//				fmt.Println("k:", k, "v:", v.Open())
//			}
//		}
//	}
//}
