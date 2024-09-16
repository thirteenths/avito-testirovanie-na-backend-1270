package mapper

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func MakeGetTenderByFilter(dom []domain.Tender) *response.GetTendersByFilter {
	tenders := make([]response.Tender, 0)

	for _, t := range dom {
		tender := response.Tender{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Status:      t.Status,
			ServiceType: t.ServiceType,
			Verstion:    t.Version,
			CreatedAt:   t.CreatedAt.String(),
		}

		tenders = append(tenders, tender)
	}

	return &response.GetTendersByFilter{
		Tenders: tenders,
	}
}
