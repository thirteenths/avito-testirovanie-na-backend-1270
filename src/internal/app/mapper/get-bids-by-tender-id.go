package mapper

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func MakeGetBidsByTenderId(dom []domain.Bid) *response.GetBidsByTenderId {
	var bids []response.Bid
	for _, b := range dom {
		var bid = response.Bid{
			ID:         b.ID,
			Name:       b.Name,
			Status:     b.Status,
			AuthorType: b.AuthorType,
			AuthorId:   b.AuthorId,
			Version:    b.Version,
			CreatedAt:  b.CreatedAt.String(),
		}

		bids = append(bids, bid)
	}
	return &response.GetBidsByTenderId{
		Bids: bids,
	}
}
