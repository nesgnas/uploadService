package main

import (
	"github.com/gin-gonic/gin"
	"uploadService/cloudbucket"
	"uploadService/middleware"
)

func main() {

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", func(c *gin.Context) {

		file, FileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": err,
				"file": FileHeader, "filee": file})
		}
		c.JSON(200, gin.H{
			"file":  FileHeader,
			"filee": file})

	})

	r.POST("/cloud-storage", cloudbucket.HandleFileUploadToBucket)

	//r.GET("/cloud-storage", cloudbucket.GenerateV4URL)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
