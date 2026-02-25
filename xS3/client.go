package xS3

import (
	"crypto/tls"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	S3Client *s3.Client
}

func NewS3Client(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, usePathStyle bool, skipTLSVerify bool) (client *Client) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLSVerify},
	}
	httpClient := http.Client{Transport: tr}

	return &Client{
		S3Client: s3.New(s3.Options{
			BaseEndpoint: &endpoint,
			Region:       region,
			Credentials:  credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, stsToken),
			UsePathStyle: usePathStyle,
			HTTPClient:   &httpClient,
		}),
	}
}
