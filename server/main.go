package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	}).POST("/user/:name/*sex", func(context *gin.Context) {
		name := context.Param("name")
		sex := context.Param("sex")
		context.String(http.StatusOK, "user{%s} maybe %s", name, sex)
	}).GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		sex := c.DefaultQuery("sex", "male")
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"sex": sex,
		})
	}).POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	}).POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	}).GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	}).GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}
	/*上传文件*/
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})
	r.Run(":9000")
}
