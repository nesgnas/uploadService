package main

import (
	"github.com/gin-gonic/gin"
	"uploadService/cloudbucket"
	"uploadService/middleware"
)

func main() {
	//fmt.Println("Hello World")
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//
	//credentialsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	//if credentialsPath == "" {
	//	fmt.Println("GOOGLE_APPLICATION_CREDENTIALS environment variable is not set")
	//	return
	//}
	//
	//// Print the path for debugging purposes
	////fmt.Printf("GOOGLE_APPLICATION_CREDENTIALS is set to: %s\n", credentialsPath)
	//
	//// Read the file
	//jsonKey, err := os.ReadFile(credentialsPath)
	//if err != nil {
	//	fmt.Printf("Error reading file: %v\n", err)
	//	return
	//}
	//
	//// Print the content for verification (optional)
	////fmt.Println("File read successfully")
	//fmt.Println(string(jsonKey))

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cloud-storage", cloudbucket.HandleFileUploadToBucket)

	r.GET("/cloud-storage", cloudbucket.GenerateV4URL)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
