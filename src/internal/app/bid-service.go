package app

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app/consts"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app/mapper"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/request"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"
)

type BidService struct {
	storage storage.Storage
}

func NewBidService(storage storage.Storage) *BidService {
	return &BidService{
		storage: storage,
	}
}

func (s *BidService) CreateBid(req request.CreateBid) (response.CreateBid, error) {
	var exist bool
	var err error

	if req.AuthorType == "User" {
		exist, err = s.storage.CheckUserIsExistById(req.AuthorId)
		if err != nil {
			return response.CreateBid{}, err
		}
	} else if req.AuthorType == "Organization" {
		exist, err = s.storage.CheckOrganizationIsExist(req.AuthorId)
		if err != nil {
			return response.CreateBid{}, err
		}
	}

	if !exist {
		return response.CreateBid{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckTenderIsExist(req.TenderId)
	if err != nil {
		return response.CreateBid{}, err
	}

	if !exist {
		return response.CreateBid{}, consts.TenderIsNotExistError
	}

	res, err := s.storage.CreateBid(*mapper.ParseCreateBid(req))
	if err != nil {
		return response.CreateBid{}, err
	}

	return *mapper.MakeCreateBid(res), nil
}

func (s *BidService) GetBidsByUsername(filter request.GetBidsByUsernameFilter) (response.GetBidsByUsername, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.GetBidsByUsername{}, err
	}

	if !exist {
		return response.GetBidsByUsername{}, consts.UserIsNotExistError
	}

	res, err := s.storage.GetBidsByFilter(filter.Limit, filter.Offset, filter.Username)
	if err != nil {
		return response.GetBidsByUsername{}, err
	}

	return *mapper.MakeGetBidsByUsername(res), nil
}

func (s *BidService) GetBidsByTenderId(filter request.GetBidsByTenderIDFilter) (response.GetBidsByTenderId, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.GetBidsByTenderId{}, err
	}

	if !exist {
		return response.GetBidsByTenderId{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckTenderIsExist(filter.TenderId)
	if err != nil {
		return response.GetBidsByTenderId{}, err
	}

	if !exist {
		return response.GetBidsByTenderId{}, consts.TenderIsNotExistError
	}

	exist, err = s.storage.CheckUserTender(filter.Username, filter.TenderId)
	if err != nil {
		return response.GetBidsByTenderId{}, err
	}

	if !exist {
		return response.GetBidsByTenderId{}, consts.UserHasNoRights
	}

	res, err := s.storage.GetBidsByTenderIdByFilter(filter.Limit, filter.Offset, filter.Username, filter.TenderId)
	if err != nil {
		return response.GetBidsByTenderId{}, err
	}

	return *mapper.MakeGetBidsByTenderId(res), nil
}

func (s *BidService) GetBidStatusById(filter request.GetBidStatusById) (string, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", consts.BidIsNotExistError
	}

	res, err := s.storage.GetBidStatusById(filter.BidId)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *BidService) UpdateBidStatusById(filter request.UpdateBidStatusById) (response.UpdateBidStatusById, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateBidStatusById{}, err
	}

	if !exist {
		return response.UpdateBidStatusById{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return response.UpdateBidStatusById{}, err
	}

	if !exist {
		return response.UpdateBidStatusById{}, consts.BidIsNotExistError
	}

	err = s.storage.UpdateBidStatus(filter.BidId, filter.Status)
	if err != nil {
		return response.UpdateBidStatusById{}, err
	}

	res, err := s.storage.GetBidsById(filter.BidId)
	if err != nil {
		return response.UpdateBidStatusById{}, err
	}

	return *mapper.MakeUpdateBidStatusByID(res), nil
}

func (s *BidService) UpdateBidParams(filter request.UpdateBidParamsFilter, req request.UpdateBidParams) (response.UpdateBidParams, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateBidParams{}, err
	}

	if !exist {
		return response.UpdateBidParams{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return response.UpdateBidParams{}, err
	}

	if !exist {
		return response.UpdateBidParams{}, consts.BidIsNotExistError
	}

	res, err := s.storage.GetBidsById(filter.BidId)
	if err != nil {
		return response.UpdateBidParams{}, err
	}

	if req.Name != "" {
		res.Name = req.Name
	}

	if req.Description != "" {
		res.Description = req.Description
	}

	res.Version += 1

	err = s.storage.UpdateBidById(filter.BidId, res)
	if err != nil {
		return response.UpdateBidParams{}, err
	}

	res, err = s.storage.GetBidsById(filter.BidId)
	if err != nil {
		return response.UpdateBidParams{}, err
	}

	return *mapper.MakeUpdateBidParams(res), nil
}

func (s *BidService) UpdateBidDecision(filter request.UpdateBidDecisionFilter) (response.UpdateBidDecision, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateBidDecision{}, err
	}

	if !exist {
		return response.UpdateBidDecision{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return response.UpdateBidDecision{}, err
	}

	if !exist {
		return response.UpdateBidDecision{}, consts.BidIsNotExistError
	}

	err = s.storage.UpdateBidDecisionById(filter.BidId, filter.Decision)
	if err != nil {
		return response.UpdateBidDecision{}, err
	}

	res, err := s.storage.GetBidsById(filter.BidId)
	if err != nil {
		return response.UpdateBidDecision{}, err
	}

	return *mapper.MakeUpdateBidDecision(res), nil
}

func (s *BidService) UpdateBidFeedbackById(filter request.UpdateBidFeedbackById) (response.UpdateBidFeedbackById, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateBidFeedbackById{}, err
	}

	if !exist {
		return response.UpdateBidFeedbackById{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return response.UpdateBidFeedbackById{}, err
	}

	if !exist {
		return response.UpdateBidFeedbackById{}, consts.BidIsNotExistError
	}

	err = s.storage.UpdateBidFeedbackById(filter.BidId, filter.BidFeedback)
	if err != nil {
		return response.UpdateBidFeedbackById{}, err
	}

	res, err := s.storage.GetBidsById(filter.BidId)
	if err != nil {
		return response.UpdateBidFeedbackById{}, err
	}

	return *mapper.MakeUpdateBidFeedbackById(res), nil
}

func (s *BidService) UpdateBidVersionRollback(filter request.UpdateBidVersionRollback) (response.UpdateBidVersionRollback, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateBidVersionRollback{}, err
	}

	if !exist {
		return response.UpdateBidVersionRollback{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckBidIsExist(filter.BidId)
	if err != nil {
		return response.UpdateBidVersionRollback{}, err
	}

	if !exist {
		return response.UpdateBidVersionRollback{}, consts.BidIsNotExistError
	}

	res, err := s.storage.GetBidVersionById(filter.BidId, filter.Version)
	if err != nil {
		return response.UpdateBidVersionRollback{}, err
	}

	err = s.storage.UpdateBidById(filter.BidId, res)
	if err != nil {
		return response.UpdateBidVersionRollback{}, err
	}

	res, err = s.storage.GetBidsById(res.ID)
	if err != nil {
		return response.UpdateBidVersionRollback{}, err
	}

	return *mapper.MakeUpdateBidVersionRollback(res), nil
}
