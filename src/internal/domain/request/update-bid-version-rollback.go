package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UpdateBidVersionRollback struct {
	BidId    string
	Version  int
	Username string
}

func NewUpdateBidVersionRollback() *UpdateBidVersionRollback {
	return &UpdateBidVersionRollback{}
}

func (u *UpdateBidVersionRollback) Bind(r *http.Request) (err error) {
	values := r.URL.Query()

	u.BidId = chi.URLParam(r, "bidId")
	if u.BidId == "" {
		return errors.New("missing bidId")
	}

	u.Version, err = strconv.Atoi(values.Get("version"))
	if err != nil {
		return errors.New("invalid version")
	}

	if values.Has("username") {
		u.Username = values.Get("username")
	} else {
		return errors.New("username is required")
	}

	return nil
}
