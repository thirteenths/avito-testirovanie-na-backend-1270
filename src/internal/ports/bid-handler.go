package ports

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/app/consts"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/request"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725744525-team-79206/zadanie-6105/internal/domain/response"
)

type BidHandler struct {
	bid app.BidService
}

func NewBidHandler(bid app.BidService) *BidHandler {
	return &BidHandler{
		bid: bid,
	}
}

func (h *BidHandler) CreateBid(w http.ResponseWriter, r *http.Request) {
	var req request.CreateBid

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.CreateBid(req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Пользователя не существует")))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Пользователь нет прав ")))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка создания тендера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) GetBidsByUsername(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidsByUsernameFilter()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidsByUsername(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) GetBidsByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidsByTenderIDFilter()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidsByTenderId(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) GetBidsStatusByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidStatusById()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) UpdateBidStatusByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidStatusById()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) UpdateBidParamsByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidParamsFilter()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var req request.UpdateBidParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidParams(*filters, req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) UpdateBidDecisionByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidDecisionFilter()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidDecision(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) UpdateBidFeedBackById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidFeedbackById()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidFeedbackById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) UpdateBidVersionRollback(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidVersionRollback()

	w.Header().Set("Content-Type", "application/json")

	if err := filters.Bind(r); err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка запроса")))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidVersionRollback(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует", filters.Username)))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf(
			"У пользователя %s нет прав", filters.Username)))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	if err != nil {
		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Ошибка сервера")))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BidHandler) GetBidReviewsById(w http.ResponseWriter, r *http.Request) {}
