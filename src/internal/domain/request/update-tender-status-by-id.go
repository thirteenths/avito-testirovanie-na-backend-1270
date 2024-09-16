package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateTenderStatusById struct {
	TenderId string `json:"tenderId"`
	Status   string `json:"status"`
	Username string `json:"username"`
}

func NewUpdateTenderStatusById() *UpdateTenderStatusById {
	return &UpdateTenderStatusById{}
}

func (u *UpdateTenderStatusById) Bind(r *http.Request) error {
	u.TenderId = chi.URLParam(r, "tenderId")
	if u.TenderId == "" {
		return errors.New("missing tenderId")
	}

	if r.URL.Query().Has("username") {
		u.Username = r.URL.Query().Get("username")
	} else {
		return errors.New("missing username")
	}

	if r.URL.Query().Has("status") {
		u.Status = r.URL.Query().Get("status")
		if u.Status != "Created" && u.Status != "Published" && u.Status != "Closed" {
			return errors.New("invalid status")
		}
	} else {
		return errors.New("missing status")
	}

	return nil
}
