package storage

import "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"

type BidStorage interface {
	CheckBidIsExist(bidId string) (bool, error)
	CheckOrganizationIsExist(organizationId string) (bool, error)
	CreateBid(bid domain.Bid) (domain.Bid, error)
	GetBidsByFilter(limit, offset int, username string) ([]domain.Bid, error)
	GetBidsByTenderIdByFilter(limit, offset int, username string, tenderId string) ([]domain.Bid, error)
	GetBidsById(bidId string) (domain.Bid, error)
	GetBidStatusById(bidId string) (string, error)
	UpdateBidStatus(bidId string, status string) error
	UpdateBidById(bidId string, bid domain.Bid) error
	UpdateBidDecisionById(bidId string, decision string) error
	UpdateBidFeedbackById(bidId string, feedback string) error
	GetBidVersionById(bidId string, version int) (domain.Bid, error)
	// GetBidReviewByTenderId(tenderId string) ([]domain.Bid, error)
}
