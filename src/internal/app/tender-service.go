package app

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app/consts"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app/mapper"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/request"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/storage"
)

type TenderService struct {
	storage storage.Storage
}

func NewTenderService(storage storage.Storage) *TenderService {
	return &TenderService{storage}
}

func (s *TenderService) CreateTender(req request.CreateTender) (res response.CreateTender, err error) {
	exist, err := s.storage.CheckUserIsExistByUsername(req.CreatorUsername)
	if err != nil {
		return response.CreateTender{}, err
	}

	if !exist {
		return response.CreateTender{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckUserOrganization(req.CreatorUsername, req.OrganizationId)
	if err != nil {
		return response.CreateTender{}, err
	}

	if !exist {
		return response.CreateTender{}, consts.UserHasNoRights
	}

	tender, err := s.storage.CreateTender(*mapper.ParseCreateTender(req))
	if err != nil {
		return response.CreateTender{}, err
	}

	return *mapper.MakeCreateTender(tender), nil
}

func (s *TenderService) GetTenderByFilter(filter request.GetTendersFilters) (response.GetTendersByFilter, error) {
	var tenders []domain.Tender
	var err error

	if len(filter.ServiceType) == 1 {
		tenders, err = s.storage.GetTendersByOneServiceType(filter.Limit, filter.Offset, filter.ServiceType)
		if err != nil {
			return response.GetTendersByFilter{}, err
		}
	} else if len(filter.ServiceType) == 2 {
		tenders, err = s.storage.GetTendersByTwoServiceType(filter.Limit, filter.Offset, filter.ServiceType)
		if err != nil {
			return response.GetTendersByFilter{}, err
		}
	} else {
		tenders, err = s.storage.GetTendersByFilter(filter.Limit, filter.Offset)
		if err != nil {
			return response.GetTendersByFilter{}, err
		}
	}

	return *mapper.MakeGetTenderByFilter(tenders), nil
}

func (s *TenderService) GetTenderByUsername(filter request.GetTendersByUsername) (response.GetTenderByUsername, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.GetTenderByUsername{}, err
	}

	if !exist {
		return response.GetTenderByUsername{}, consts.UserIsNotExistError
	}

	res, err := s.storage.GetTenderByUsername(filter.Username, filter.Limit, filter.Offset)
	if err != nil {
		return response.GetTenderByUsername{}, err
	}

	return *mapper.MakeGetTenderByUsername(res), nil
}

func (s *TenderService) GetTenderStatusById(filter request.GetTenderStatusById) (string, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckTenderIsExist(filter.TenderId)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", consts.TenderIsNotExistError
	}

	exist, err = s.storage.CheckUserTender(filter.Username, filter.TenderId)
	if err != nil {
		return "", consts.UserHasNoRights
	}

	if !exist {
		return "", consts.UserHasNoRights
	}

	res, err := s.storage.GetStatusTenderById(filter.TenderId)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *TenderService) UpdateTenderStatusById(filter request.UpdateTenderStatusById) (response.UpdateTenderStatusById, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateTenderStatusById{}, err
	}

	if !exist {
		return response.UpdateTenderStatusById{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckTenderIsExist(filter.TenderId)
	if err != nil {
		return response.UpdateTenderStatusById{}, err
	}

	if !exist {
		return response.UpdateTenderStatusById{}, consts.TenderIsNotExistError
	}

	exist, err = s.storage.CheckUserTender(filter.Username, filter.TenderId)
	if err != nil {
		return response.UpdateTenderStatusById{}, consts.UserHasNoRights
	}

	err = s.storage.UpdateStatusTenderById(filter.TenderId, filter.Status)
	if err != nil {
		return response.UpdateTenderStatusById{}, err
	}

	res, err := s.storage.GetTenderById(filter.TenderId)
	if err != nil {
		return response.UpdateTenderStatusById{}, err
	}

	return *mapper.MakeUpdateTenderStatusByID(res), nil
}

func (s *TenderService) UpdateTenderParams(filter request.UpdateTenderParamsFilter, req request.UpdateTenderParams) (response.UpdateTenderParams, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	if !exist {
		return response.UpdateTenderParams{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckTenderIsExist(filter.TenderId)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	if !exist {
		return response.UpdateTenderParams{}, consts.TenderIsNotExistError
	}

	exist, err = s.storage.CheckUserTender(filter.Username, filter.TenderId)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	if !exist {
		return response.UpdateTenderParams{}, consts.UserHasNoRights
	}

	res, err := s.storage.GetTenderById(filter.TenderId)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	// todo не делать updaye если запрос пустой

	if req.Name != "" {
		res.Name = req.Name
	}

	if req.Description != "" {
		res.Description = req.Description
	}

	if req.ServiceType != "" {
		res.ServiceType = req.ServiceType
	}

	res.Version += 1

	err = s.storage.UpdateTenderById(res)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	res, err = s.storage.GetTenderById(filter.TenderId)
	if err != nil {
		return response.UpdateTenderParams{}, err
	}

	return *mapper.MakeUpdateTenderParamsByID(res), nil
}

func (s *TenderService) UpdateTenderVersionRollback(filter request.UpdateTenderVersionRollbackFilter) (response.UpdateTenderVersionRollback, error) {
	exist, err := s.storage.CheckUserIsExistByUsername(filter.Username)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	if !exist {
		return response.UpdateTenderVersionRollback{}, consts.UserIsNotExistError
	}

	exist, err = s.storage.CheckUserTender(filter.Username, filter.TenderId)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	exist, err = s.storage.CheckVersionTenderIsExist(filter.Version, filter.TenderId)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	if !exist {
		return response.UpdateTenderVersionRollback{}, consts.VersionIsNotExistError
	}

	version, err := s.storage.GetTenderLastVersion(filter.TenderId)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	tender, err := s.storage.GetTenderVersion(filter.TenderId, filter.Version)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, consts.VersionIsNotExistError
	}

	tender.Version = version + 1
	tender.Id = filter.TenderId

	err = s.storage.UpdateTenderById(tender)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	res, err := s.storage.GetTenderById(filter.TenderId)
	if err != nil {
		return response.UpdateTenderVersionRollback{}, err
	}

	return *mapper.MakeUpdateTenderVersionRollback(res), nil
}
