package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//文件上传
func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.String(http.StatusOK,fmt.Sprintf(" '%s' uploadSussed",file.Filename))
	})
}
//中间件
func main2() {
	r := gin.Default()
	r.GET("/get",middleware1, func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname","guest")
		lastName := c.Query("lastname")

		c.String(http.StatusOK,firstName + "+" + lastName)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("msg")
		nick := c.DefaultPostForm("nick","小米")

		c.JSON(http.StatusOK,gin.H{
			"states" : "ok",
			"meassge" : message,
			"nick" : nick,
		})
	})

	r.Run()
}

func middleware1(c *gin.Context) {
	log.Println("exec middleware1")
	c.Next()
}

func main1() {
	router := gin.Default()
	router.Any("/", func(context *gin.Context) {
		context.String(http.StatusOK,"今晚约吗？")
	})

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"hello %s",name)
	})

	router.GET("user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		msg := name + "is" + action
		c.String(http.StatusOK,msg)
	})

	router.Run(":9527")
}
