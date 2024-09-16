package request

type CreateTender struct {
	Name            string `json:"name"`            // "name": "string",
	Description     string `json:"description"`     // "description": "string",
	ServiceType     string `json:"serviceType"`     // "serviceType": "Construction",
	OrganizationId  string `json:"organizationId"`  // "organizationId": "550e8400-e29b-41d4-a716-446655440000",
	CreatorUsername string `json:"creatorUsername"` // "creatorUsername": "test_user"
}
