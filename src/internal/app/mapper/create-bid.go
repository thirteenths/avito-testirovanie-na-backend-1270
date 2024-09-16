package mapper

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/request"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func ParseCreateBid(bid request.CreateBid) *domain.Bid {
	return &domain.Bid{
		Name:        bid.Name,
		Description: bid.Description,
		TenderId:    bid.TenderId,
		AuthorId:    bid.AuthorId,
		AuthorType:  bid.AuthorType,
	}
}

func MakeCreateBid(dom domain.Bid) *response.CreateBid {
	return &response.CreateBid{
		ID:         dom.ID,
		Name:       dom.Name,
		Status:     dom.Status,
		AuthorType: dom.AuthorType,
		AuthorId:   dom.AuthorId,
		Version:    dom.Version,
		CreatedAt:  dom.CreatedAt.String(),
	}
}
