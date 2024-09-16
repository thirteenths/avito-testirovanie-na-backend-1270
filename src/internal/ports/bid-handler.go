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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.CreateBid(req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating bid: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating bid: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь нет прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error creating bid: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(" Тендер не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error creating bid: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка создания предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка создания предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) GetBidsByUsername(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidsByUsernameFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidsByUsername(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error get bids by username: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if err != nil {
		log.Printf("Error get bids by username: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка создания предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) GetBidsByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidsByTenderIDFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidsByTenderId(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error get bids by tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error get bids by tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error get bids by tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Тендер не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error get bids by tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка получения предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) GetBidsStatusById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetBidStatusById()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.GetBidStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error get bid by status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error get bid by status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.BidIsNotExistError) {
		log.Printf("Error get bid by status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Предложения не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error get bid by status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка получения предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) UpdateBidStatusById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidStatusById()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error update bid status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error update bid status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.BidIsNotExistError) {
		log.Printf("Error update bid status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Предложения не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error update bid status: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка получения предложения")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) UpdateBidParamsByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidParamsFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	var req request.UpdateBidParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidParams(*filters, req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error update bid params: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error update bid params: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.BidIsNotExistError) {
		log.Printf("Error update bid params: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Предложение не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error update bid params: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) UpdateBidDecisionByTenderId(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidDecisionFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidDecision(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error update bid decision: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error update bid decision: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.BidIsNotExistError) {
		log.Printf("Error update bid decision: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Предложение не существует")))

		http.Error(w, string(b), http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error update bid decision: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) UpdateBidFeedBackById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidFeedbackById()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidFeedbackById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error update bid feedback: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error update bid feedback: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if err != nil {
		log.Printf("Error update bid feedback: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) UpdateBidVersionRollback(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateBidVersionRollback()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error bind request: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.bid.UpdateBidVersionRollback(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error update bid version: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error update bid version: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователя не прав")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if err != nil {
		log.Printf("Error update bid version: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *BidHandler) GetBidReviewsById(w http.ResponseWriter, r *http.Request) {}
