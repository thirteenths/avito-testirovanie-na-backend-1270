package mapper

import (
	"github.com/google/uuid"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/request"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func MakeCreateTender(dom domain.Tender) *response.CreateTender {
	return &response.CreateTender{
		Id:          dom.Id,
		Name:        dom.Name,
		Description: dom.Description,
		Status:      dom.Status,
		ServiceType: dom.ServiceType,
		Version:     dom.Version,
		CreatedAt:   dom.CreatedAt.String(),
	}
}

func ParseCreateTender(tender request.CreateTender) *domain.Tender {
	return &domain.Tender{
		Id:             uuid.New().String(),
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		OrganizationId: tender.OrganizationId,
	}
}
