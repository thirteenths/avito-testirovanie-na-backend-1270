package storage

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
)

type TenderStorage interface {
	CreateTender(tender domain.Tender) (domain.Tender, error)
	GetAllTenders() ([]domain.Tender, error)
	GetTendersByFilter(limit, offset int) ([]domain.Tender, error)
	GetTendersByOneServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error)
	GetTendersByTwoServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error)
	GetTenderById(id string) (domain.Tender, error)
	GetTenderByUsername(username string, limit, offset int) ([]domain.Tender, error)
	GetStatusTenderById(id string) (string, error)
	UpdateStatusTenderById(tenderId string, status string) error
	UpdateTenderById(tender domain.Tender) error
	GetTenderVersion(tenderId string, version int) (domain.Tender, error)
	CheckTenderIsExist(tenderId string) (bool, error)
	CheckVersionTenderIsExist(version int, tenderId string) (bool, error)
	GetTenderLastVersion(tenderId string) (int, error)
}
