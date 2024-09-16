package domain

import "time"

type Tender struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ServiceType    string    `json:"service_type"`
	Status         string    `json:"status"`
	OrganizationId string    `json:"organization_id"`
	Version        int       `json:"version"`
	VersionId      string    `json:"version_id"`
	CreatedAt      time.Time `json:"created_at"`
}
