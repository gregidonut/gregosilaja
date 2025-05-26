package models

import "time"

type Blog struct {
	ID           string    `json:"id"`
	Path         string    `json:"path"`
	Body         string    `json:"body"`
	LastModified time.Time `json:"lastModified"`
}
