package models

import "time"

type BlogAsset struct {
	Ext       string `json:"ext"`
	W         int    `json:"w"`
	H         int    `json:"h"`
	B64String string `json:"b64String"`
}

type Blog struct {
	ID           string               `json:"id"`
	Path         string               `json:"path"`
	Body         string               `json:"body"`
	LastModified time.Time            `json:"lastModified"`
	Assets       map[string]BlogAsset `json:"assets"`
}
