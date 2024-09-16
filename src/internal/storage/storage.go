package storage

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
)

type Storage interface {
	TenderStorage
	EmployeeStorage
	BidStorage
}

type storage struct {
	Tender   TenderStorage
	Employee EmployeeStorage
	Bid      BidStorage
}

func (s *storage) CheckUserBid(username string, bidId string) (bool, error) {
	return s.Employee.CheckUserBid(username, bidId)
}

func (s *storage) CheckVersionTenderIsExist(version int, tenderId string) (bool, error) {
	return s.Tender.CheckVersionTenderIsExist(version, tenderId)
}

func (s *storage) CheckBidIsExist(bidId string) (bool, error) {
	return s.Bid.CheckBidIsExist(bidId)
}

func (s *storage) CreateBid(bid domain.Bid) (domain.Bid, error) {
	return s.Bid.CreateBid(bid)
}

func (s *storage) GetBidsByFilter(limit, offset int, username string) ([]domain.Bid, error) {
	return s.Bid.GetBidsByFilter(limit, offset, username)
}

func (s *storage) GetBidsByTenderIdByFilter(limit, offset int, username string, tenderId string) ([]domain.Bid, error) {
	return s.Bid.GetBidsByTenderIdByFilter(limit, offset, username, tenderId)
}

func (s *storage) GetBidsById(bidId string) (domain.Bid, error) {
	return s.Bid.GetBidsById(bidId)
}

func (s *storage) GetBidStatusById(bidId string) (string, error) {
	return s.Bid.GetBidStatusById(bidId)
}

func (s *storage) UpdateBidStatus(bidId string, status string) error {
	return s.Bid.UpdateBidStatus(bidId, status)
}

func (s *storage) UpdateBidById(bidId string, bid domain.Bid) error {
	return s.Bid.UpdateBidById(bidId, bid)
}

func (s *storage) UpdateBidDecisionById(bidId string, decision string) error {
	return s.Bid.UpdateBidDecisionById(bidId, decision)
}

func (s *storage) UpdateBidFeedbackById(bidId string, feedback string) error {
	return s.Bid.UpdateBidFeedbackById(bidId, feedback)
}

func (s *storage) GetBidVersionById(bidId string, version int) (domain.Bid, error) {
	return s.Bid.GetBidVersionById(bidId, version)
}

func (s *storage) CreateTender(tender domain.Tender) (domain.Tender, error) {
	return s.Tender.CreateTender(tender)
}

func (s *storage) GetAllTenders() ([]domain.Tender, error) {
	return s.Tender.GetAllTenders()
}

func (s *storage) GetTendersByFilter(limit, offset int) ([]domain.Tender, error) {
	return s.Tender.GetTendersByFilter(limit, offset)
}

func (s *storage) GetTendersByOneServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error) {
	return s.Tender.GetTendersByOneServiceType(limit, offset, serviceType)
}

func (s *storage) GetTendersByTwoServiceType(limit, offset int, serviceType []string) ([]domain.Tender, error) {
	return s.Tender.GetTendersByTwoServiceType(limit, offset, serviceType)
}

func (s *storage) GetTenderById(id string) (domain.Tender, error) {
	return s.Tender.GetTenderById(id)
}

func (s *storage) GetStatusTenderById(id string) (string, error) {
	return s.Tender.GetStatusTenderById(id)
}

func (s *storage) UpdateStatusTenderById(tenderId string, status string) error {
	return s.Tender.UpdateStatusTenderById(tenderId, status)
}

func (s *storage) UpdateTenderById(tender domain.Tender) error {
	return s.Tender.UpdateTenderById(tender)
}

func (s *storage) CheckUserIsExistByUsername(username string) (bool, error) {
	return s.Employee.CheckUserIsExistByUsername(username)
}

func (s *storage) CheckUserOrganization(username string, organizationId string) (bool, error) {
	return s.Employee.CheckUserOrganization(username, organizationId)
}

func (s *storage) GetTenderByUsername(username string, limit, offset int) ([]domain.Tender, error) {
	return s.Tender.GetTenderByUsername(username, limit, offset)
}

func (s *storage) CheckUserTender(username string, tenderId string) (bool, error) {
	return s.Employee.CheckUserTender(username, tenderId)
}

func (s *storage) GetTenderVersion(tenderId string, version int) (domain.Tender, error) {
	return s.Tender.GetTenderVersion(tenderId, version)
}

func (s *storage) CheckUserIsExistById(userId string) (bool, error) {
	return s.Employee.CheckUserIsExistById(userId)
}

func (s *storage) CheckTenderIsExist(tenderId string) (bool, error) {
	return s.Tender.CheckTenderIsExist(tenderId)
}

func (s *storage) GetTenderLastVersion(tenderId string) (int, error) {
	return s.Tender.GetTenderLastVersion(tenderId)
}

func (s *storage) CheckOrganizationIsExist(organizationId string) (bool, error) {
	return s.Bid.CheckOrganizationIsExist(organizationId)
}

func NewStorage(tender TenderStorage, employee EmployeeStorage, bid BidStorage) Storage {
	return &storage{
		Tender:   tender,
		Employee: employee,
		Bid:      bid,
	}
}
