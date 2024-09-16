package request

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateBidParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBidParamsFilter struct {
	BidId    string
	Username string
}

func NewUpdateBidParamsFilter() *UpdateBidParamsFilter {
	return &UpdateBidParamsFilter{}
}

func (u *UpdateBidParamsFilter) Bind(r *http.Request) error {
	values := r.URL.Query()

	u.BidId = chi.URLParam(r, "bidId")
	if u.BidId == "" {
		return errors.New("missing bidId")
	}

	if values.Has("username") {
		u.Username = values.Get("username")
	} else {
		return errors.New("username is required")
	}

	return nil
}
