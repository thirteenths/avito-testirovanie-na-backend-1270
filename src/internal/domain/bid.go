package domain

import "time"

type Bid struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	TenderId    string    `json:"tender_id"`
	AuthorType  string    `json:"author_type"`
	AuthorId    string    `json:"author_id"`
	Version     int       `json:"version"`
	VersionId   string    `json:"version_id"`
	CreatedAt   time.Time `json:"created_at"`
}
