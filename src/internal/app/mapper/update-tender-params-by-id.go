package mapper

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

func MakeUpdateTenderParamsByID(dom domain.Tender) *response.UpdateTenderParams {
	return &response.UpdateTenderParams{
		Id:          dom.Id,
		Name:        dom.Name,
		Description: dom.Description,
		Status:      dom.Status,
		ServiceType: dom.ServiceType,
		Verstion:    dom.Version,
		CreatedAt:   dom.CreatedAt.String(),
	}
}
