package model

import "time"

type UrlInput struct {
	Url      *string    `json:"url"`
	ExpireAt *time.Time `json:"expireAt"`
}
