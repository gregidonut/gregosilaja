package main

import (
	"cmp"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/models/blog"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/s3helpers"
	"slices"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/utils"
	"github.com/sst/sst/v3/sdk/golang/resource"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	blogName, err := resource.Get("Blog", "name")
	if err != nil {
		return utils.APIServerError(err)
	}

	bucket, err := s3helpers.NewBucket(ctx, blogName.(string))
	if err != nil {
		return utils.APIServerError(err)
	}

	objectList, err := bucket.ListObjects("blogs/")
	if err != nil {
		return utils.APIServerError(err)
	}

	slices.SortFunc(objectList, func(a, b types.Object) int {
		return cmp.Compare(*a.Key, *b.Key)
	})

	blogs, err := blog.NewBlogs(objectList, bucket)
	if err != nil {
		return utils.APIServerError(err)
	}

	body, _ := json.MarshalIndent(blogs, "", "    ")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(handler)
}
