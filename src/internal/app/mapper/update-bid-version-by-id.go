package mapper

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func MakeUpdateBidVersionRollback(dom domain.Bid) *response.UpdateBidVersionRollback {
	return &response.UpdateBidVersionRollback{
		ID:         dom.ID,
		Name:       dom.Name,
		Status:     dom.Status,
		AuthorType: dom.AuthorType,
		AuthorId:   dom.AuthorId,
		Version:    dom.Version,
		CreatedAt:  dom.CreatedAt.String(),
	}
}
