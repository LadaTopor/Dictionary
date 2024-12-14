package service

import "time"

type Word struct {
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

type Reports struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
