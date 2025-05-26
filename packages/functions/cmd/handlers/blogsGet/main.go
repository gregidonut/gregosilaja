package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/models"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/s3helpers"
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

	objectList, err := bucket.ListObjects()
	if err != nil {
		return utils.APIServerError(err)
	}

	var blogs []models.Blog
	for _, obj := range objectList {
		if obj.Key == nil {
			continue
		}

		content, err := bucket.GetObjectContents(*obj.Key)
		if err != nil {
			return utils.APIServerError(err)
		}

		blogs = append(blogs, models.Blog{
			ID:           *obj.Key,
			Body:         content,
			LastModified: *obj.LastModified,
		})
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
