package model

import "time"

type Book struct {
	ID        int64
	Title     string
	Author    string
	Publisher string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
