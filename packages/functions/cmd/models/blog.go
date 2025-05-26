package models

import "time"

type Blog struct {
	ID           string    `json:"id"`
	Body         string    `json:"body"`
	LastModified time.Time `json:"lastModified"`
}
