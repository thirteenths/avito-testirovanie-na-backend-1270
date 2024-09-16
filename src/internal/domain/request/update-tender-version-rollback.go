package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UpdateTenderVersionRollbackFilter struct {
	TenderId string
	Version  int
	Username string
}

func NewUpdateTenderVersionRollbackFilter() *UpdateTenderVersionRollbackFilter {
	return &UpdateTenderVersionRollbackFilter{}
}

func (u *UpdateTenderVersionRollbackFilter) Bind(r *http.Request) error {
	var err error

	u.TenderId = chi.URLParam(r, "tenderId")
	if u.TenderId == "" {
		return errors.New("missing tenderId")
	}

	u.Version, err = strconv.Atoi(chi.URLParam(r, "version"))
	if err != nil {
		return errors.New("invalid version")
	}

	if r.URL.Query().Has("username") {
		u.Username = r.URL.Query().Get("username")
	} else {
		return errors.New("missing username")
	}

	return nil
}
