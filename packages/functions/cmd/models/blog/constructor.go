package blog

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/s3helpers"
	"github.com/gregidonut/gregosilaja/packages/functions/cmd/utils"
	"github.com/yuin/goldmark"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func NewBlog(
	obj types.Object,
	bucket *s3helpers.Bucket,
	ba map[string]*BlogAsset,
) (*Blog, error) {
	key := *obj.Key
	content, err := bucket.GetObjectContents(key)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return nil, err
	}

	id := regexp.MustCompile(`[/.]`).Split(key, -1)[1]

	return &Blog{
		ID:           id,
		Path:         key,
		Body:         buf.String(),
		LastModified: *obj.LastModified,
		Assets:       ba,
	}, nil
}

func NewBlogAsset(key string, c []byte,
) (*BlogAsset, error) {
	width, height, err := utils.GetImageDimensions(c)
	if err != nil {
		return nil, err
	}

	ext := strings.Trim(filepath.Ext(key), ".")

	switch ext {
	case "jpg", "jpeg":
		break
	default:
		return nil, fmt.Errorf("unsupported file type")
	}

	// b for bytes
	b := c
	if width > 400 {
		b, err = utils.ResConv(c)
		if err != nil {
			return nil, err
		}
		ext = "webp"
	}

	return &BlogAsset{
		Ext:       ext,
		W:         width,
		H:         height,
		B64String: base64.StdEncoding.EncodeToString(b),
	}, nil
}

func NewBlogs(objectList []types.Object,
	bucket *s3helpers.Bucket,
) ([]*Blog, error) {
	blogAssets := make(map[string]*BlogAsset)
	blogs := make([]*Blog, 0)
	for _, obj := range objectList {
		if obj.Key == nil {
			continue
		}

		key := *obj.Key
		if strings.Contains(key, "assets") {
			// c for content
			c, err := bucket.GetObjectContents(key)
			if err != nil {
				return nil, err
			}

			k := path.Join("assets/", path.Base(key))
			blogAsset, err := NewBlogAsset(k, c)
			if err != nil {
				return nil, err
			}

			blogAssets[k] = blogAsset
			continue
		}

		blogEntry, err := NewBlog(obj, bucket, blogAssets)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blogEntry)

		blogAssets = make(map[string]*BlogAsset)
	}
	return blogs, nil
}
