package xS3

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) UploadObject(bucket string, key string, data []byte) (*transfermanager.UploadObjectOutput, error) {
	return transfermanager.New(c.S3Client).UploadObject(context.TODO(), &transfermanager.UploadObjectInput{Bucket: &bucket, Key: &key, Body: bytes.NewReader(data)})
}

func (c *Client) GetObject(bucket string, key string) ([]byte, error) {
	getObjectOutput, err := transfermanager.New(c.S3Client).GetObject(context.TODO(), &transfermanager.GetObjectInput{Bucket: &bucket, Key: &key})
	if err != nil {
		return nil, err
	}
	return io.ReadAll(getObjectOutput.Body)
}

func (c *Client) DownloadObject(bucket string, key string, filename string) (err error) {
	bs, err := c.GetObject(bucket, key)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bs, 0644)
}

func (c *Client) DeleteObject(bucket string, key string) (*s3.DeleteObjectOutput, error) {
	return c.S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: &bucket, Key: &key})
}
