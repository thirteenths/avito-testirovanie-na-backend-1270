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

type TenderHandler struct {
	tender app.TenderService
}

func NewTenderHandler(tender app.TenderService) *TenderHandler {
	return &TenderHandler{
		tender: tender,
	}
}

func (h *TenderHandler) GetTenderByFilter(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetTendersByFilters()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error binding filters: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.GetTenderByFilter(*filters)
	if err != nil {
		log.Printf("Error server: %v\n", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка возврата ответа")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *TenderHandler) CreateTender(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTender

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.CreateTender(req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя %s не существует или организации  id = %s",
			req.CreatorUsername,
			req.OrganizationId,
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь %s нет прав ", req.CreatorUsername)))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка создания тендера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка возврата ответа")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *TenderHandler) GetTendersByUsername(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetTendersByUsername()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.GetTenderByUsername(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя не существует",
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка создания тендера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding response: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}

func (h *TenderHandler) GetTenderStatusById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewGetTenderStatusById()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.GetTenderStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя не существует",
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь нет прав ")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Тендер не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
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

func (h *TenderHandler) UpdateTenderStatusById(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateTenderStatusById()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.UpdateTenderStatusById(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя не существует",
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь нет прав ")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Тендер не существует")))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
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

func (h *TenderHandler) UpdateTenderParams(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateTenderParamsFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	var req request.UpdateTenderParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.UpdateTenderParams(*filters, req)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя не существует",
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь нет прав ")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Тендер не существует")))

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
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

func (h *TenderHandler) UpdateTenderVersionRollback(w http.ResponseWriter, r *http.Request) {
	filters := request.NewUpdateTenderVersionRollbackFilter()

	if err := filters.Bind(r); err != nil {
		log.Printf("Error decoding request body: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка запроса")))

		http.Error(w, string(b), http.StatusBadRequest)

		return
	}

	res, err := h.tender.UpdateTenderVersionRollback(*filters)
	if errors.Is(err, consts.UserIsNotExistError) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf(
			"Пользователя не существует",
		)))

		http.Error(w, string(b), http.StatusUnauthorized)

		return
	}

	if errors.Is(err, consts.UserHasNoRights) {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Пользователь нет прав ")))

		http.Error(w, string(b), http.StatusForbidden)

		return
	}

	if errors.Is(err, consts.TenderIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Тендер не существует")))

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if errors.Is(err, consts.VersionIsNotExistError) {
		log.Printf("Error creating tender: %v", err)

		_ = json.NewEncoder(w).Encode(response.NewError(fmt.Sprintf("Версии не существует")))

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error creating tender: %v", err)
		b, _ := json.Marshal(response.NewError(fmt.Sprintf("Ошибка сервера")))

		http.Error(w, string(b), http.StatusInternalServerError)

		return
	}
}
