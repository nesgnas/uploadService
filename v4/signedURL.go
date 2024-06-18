package v4

import (
	"fmt"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
)

// generateV4GetObjectSignedURL generates object signed URL with GET method.
func GenerateV4PutObjectSignedURL(w io.Writer, client *storage.Client, bucketName, objectName, serviceAccount string) (string, error) {
	// [START storage_generate_upload_signed_url_v4]
	jsonKey, err := ioutil.ReadFile(serviceAccount)
	if err != nil {
		return "", fmt.Errorf("cannot read the JSON key file, err: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}

	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(168 * time.Hour),
	}

	u, err := storage.SignedURL(bucketName, objectName, opts)
	if err != nil {
		return "", fmt.Errorf("Unable to generate a signed URL: %v", err)
	}

	fmt.Fprintln(w, "Generated PUT signed URL:")
	fmt.Fprintf(w, "%q\n", u)
	fmt.Fprintln(w, "You can use this URL with any user agent, for example:")
	fmt.Fprintf(w, "curl -X PUT -H 'Content-Type: application/octet-stream' --upload-file my-file %q\n", u)
	// [END storage_generate_upload_signed_url_v4]
	return u, nil
}

func GenerateV4GetObjectSignedURL(w io.Writer, client *storage.Client, bucketName, objectName, serviceAccount string) (string, error) {
	// [START storage_generate_signed_url_v4]
	jsonKey, err := ioutil.ReadFile(serviceAccount)
	if err != nil {
		return "", fmt.Errorf("cannot read the JSON key file, err: %v", err)
	}

	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(168 * time.Hour),
	}

	u, err := storage.SignedURL(bucketName, objectName, opts)
	if err != nil {
		return "", fmt.Errorf("Unable to generate a signed URL: %v", err)
	}

	fmt.Fprintln(w, "Generated GET signed URL:")
	fmt.Fprintf(w, "%q\n", u)
	fmt.Fprintln(w, "You can use this URL with any user agent, for example:")
	fmt.Fprintf(w, "curl %q\n", u)
	// [END storage_generate_signed_url_v4]
	return u, nil
}
