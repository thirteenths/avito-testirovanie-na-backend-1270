package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateTenderParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"serviceType"`
}

type UpdateTenderParamsFilter struct {
	TenderId string
	Username string
}

func NewUpdateTenderParamsFilter() *UpdateTenderParamsFilter {
	return &UpdateTenderParamsFilter{}
}

func (u *UpdateTenderParamsFilter) Bind(r *http.Request) error {
	u.TenderId = chi.URLParam(r, "tenderId")
	if u.TenderId == "" {
		return errors.New("missing tenderId")
	}

	if r.URL.Query().Has("username") {
		u.Username = r.URL.Query().Get("username")
	} else {
		return errors.New("missing username")
	}

	return nil
}
