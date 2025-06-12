package main

import (
	"bytes"
	"cmp"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/models"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/s3helpers"

	_ "image/jpeg"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/utils"
	"github.com/sst/sst/v3/sdk/golang/resource"
	"github.com/yuin/goldmark"
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

	blogAssets := map[string]models.BlogAsset{}
	var blogs []models.Blog
	for _, obj := range objectList {
		if obj.Key == nil {
			continue
		}

		key := *obj.Key
		if strings.Contains(key, "assets") {
			// c for content
			c, err := bucket.GetObjectContents(key)
			if err != nil {
				return utils.APIServerError(err)
			}

			width, height, err := utils.GetImageDimensions(c)
			if err != nil {
				return utils.APIServerError(err)
			}

			ext := strings.Trim(filepath.Ext(key), ".")

			switch ext {
			case "jpg", "jpeg":
				break
			default:
				return utils.APIServerError(fmt.Errorf("unsupported file type"))
			}

			// b for bytes
			b := c
			if width > 400 {
				b, err = utils.ResConv(c)
				if err != nil {
					return utils.APIServerError(err)
				}
				ext = "webp"
			}

			k := path.Join("assets/", path.Base(key))
			blogAssets[k] = models.BlogAsset{
				Ext:       ext,
				W:         width,
				H:         height,
				B64String: base64.StdEncoding.EncodeToString(b),
			}
			continue
		}

		content, err := bucket.GetObjectContents(key)
		if err != nil {
			return utils.APIServerError(err)
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(content, &buf); err != nil {
			return utils.APIServerError(err)
		}

		id := regexp.MustCompile(`[/.]`).Split(key, -1)[1]

		blogs = append(blogs, models.Blog{
			ID:           id,
			Path:         key,
			Body:         buf.String(),
			LastModified: *obj.LastModified,
			Assets:       blogAssets,
		})

		blogAssets = map[string]models.BlogAsset{}
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
