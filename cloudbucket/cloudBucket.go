package cloudbucket

import (
	"bytes"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	v4 "uploadService/v4"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) {
	bucket := "image-web-storage" //your bucket name

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}

	signedURL := GenerateV4URL(c, uploadedFile.Filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded successfully",

		"pathname": u.EscapedPath(),
		"signURL":  signedURL,
	})
}

func GenerateV4URL(c *gin.Context, object string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	putBuf := new(bytes.Buffer)

	bucket := "image-web-storage"
	serviceAccount := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	value, err := v4.GenerateV4GetObjectSignedURL(putBuf, storageClient, bucket, object, serviceAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return ""
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"message":  "success",
	//	"pathname": value,
	//})
	return value
}
