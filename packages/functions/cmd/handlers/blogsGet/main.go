package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/models"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/s3helpers"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/utils"
	"github.com/sst/sst/v3/sdk/golang/resource"
	"github.com/yuin/goldmark"
	"regexp"
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

	var blogs []models.Blog
	for _, obj := range objectList {
		if obj.Key == nil {
			continue
		}

		key := *obj.Key

		content, err := bucket.GetObjectContents(key)
		if err != nil {
			return utils.APIServerError(err)
		}
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(content), &buf); err != nil {
			return utils.APIServerError(err)
		}

		blogs = append(blogs, models.Blog{
			ID:           regexp.MustCompile(`[/.]`).Split(key, -1)[1],
			Path:         key,
			Body:         buf.String(),
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
