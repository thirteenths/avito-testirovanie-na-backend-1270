package domain

type OrganizationResponse struct {
	Id             int `json:"id"`
	OrganizationId int `json:"organization_id"`
	UserId         int `json:"user_id"`
}
