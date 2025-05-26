package s3helpers

import (
	"errors"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// GetObjectContents gets an object from a bucket and returns a string of it's contents.
func (b Bucket) GetObjectContents(objectKey string) (string, error) {
	result, err := b.S3Client.GetObject(b.Ctx, &s3.GetObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		var noKey *types.NoSuchKey
		if errors.As(err, &noKey) {
			log.Printf("Can't get object %s from bucket %s. No such key exists.\n", objectKey, b.Name)
			err = noKey
		} else {
			log.Printf("Couldn't get object %v:%v. Here's why: %v\n", b.Name, objectKey, err)
		}
		return "", err
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	return string(body), err
}
