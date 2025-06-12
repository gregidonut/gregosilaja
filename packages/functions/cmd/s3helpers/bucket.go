package s3helpers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/utils"
)

// Bucket encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.
type Bucket struct {
	Ctx      context.Context
	Name     string
	S3Client *s3.Client
}

func NewBucket(ctx context.Context, name string) (*Bucket, error) {
	client, err := utils.InitializeS3Client()
	if err != nil {
		return nil, err
	}

	return &Bucket{
		Ctx:      ctx,
		Name:     name,
		S3Client: client,
	}, nil
}
